// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kusciadeployment

import (
	"fmt"
	"reflect"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/encoding/protojson"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/secretflow/kuscia/pkg/common"
	kusciav1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	"github.com/secretflow/kuscia/pkg/utils/nlog"
)

const (
	configTemplateVolumeName = "config-template"
	kusciaDeploymentName     = "KusciaDeployment"
)

// ProcessKusciaDeployment processes kuscia deployment resource.
func (c *Controller) ProcessKusciaDeployment(ctx context.Context, kd *kusciav1alpha1.KusciaDeployment) (err error) {
	preKdStatus := kd.Status.DeepCopy()
	partyKitInfos, err := c.buildPartyKitInfos(kd)
	if err != nil {
		return c.handleError(ctx, preKdStatus, kd, err)
	}

	if err = c.syncResources(ctx, partyKitInfos); err != nil {
		return c.handleError(ctx, preKdStatus, kd, err)
	}

	if c.refreshPartyDeploymentStatuses(kd, partyKitInfos) {
		return c.updateKusciaDeploymentStatus(ctx, kd)
	}

	return nil
}

func (c *Controller) refreshPartyDeploymentStatuses(kd *kusciav1alpha1.KusciaDeployment, partyKitInfos map[string]*PartyKitInfo) bool {
	updated := false
	if kd.Status.TotalParties == 0 {
		updated = true
		kd.Status.TotalParties = len(partyKitInfos)
	}

	if kd.Status.PartyDeploymentStatuses == nil {
		kd.Status.PartyDeploymentStatuses = make(map[string]map[string]*kusciav1alpha1.KusciaDeploymentPartyStatus)
	}

	availableParties := 0
	hasPartialAvailableParty := false
	for _, partyKitInfo := range partyKitInfos {
		deployment, err := c.deploymentLister.Deployments(partyKitInfo.domainID).Get(partyKitInfo.dkInfo.deploymentName)
		if err != nil {
			nlog.Warnf("Failed to update party deployment %v/%v status for kuscia deployment %v, %v", partyKitInfo.domainID, partyKitInfo.dkInfo.deploymentName, kd.Name, err)
			continue
		}

		changed := refreshPartyDeploymentStatus(kd.Status.PartyDeploymentStatuses, deployment, partyKitInfo.role)
		if changed {
			updated = true
		}

		if deployment.Status.AvailableReplicas > 0 {
			availableParties++
			if deployment.Status.AvailableReplicas < deployment.Status.Replicas {
				hasPartialAvailableParty = true
			}
		}
	}

	if kd.Status.AvailableParties != availableParties {
		kd.Status.AvailableParties = availableParties
		updated = true
	}

	kdStatusPhase := kusciav1alpha1.KusciaDeploymentPhaseProgressing
	if availableParties == kd.Status.TotalParties {
		if hasPartialAvailableParty {
			kdStatusPhase = kusciav1alpha1.KusciaDeploymentPhasePartialAvailable
		} else {
			kdStatusPhase = kusciav1alpha1.KusciaDeploymentPhaseAvailable
		}
	}

	if kd.Status.Phase != kdStatusPhase {
		kd.Status.Phase = kdStatusPhase
		kd.Status.Reason = ""
		kd.Status.Message = ""
		updated = true
	}

	return updated
}

func refreshPartyDeploymentStatus(partyDeploymentStatuses map[string]map[string]*kusciav1alpha1.KusciaDeploymentPartyStatus, deployment *appsv1.Deployment, role string) bool {
	curDepStatus := &kusciav1alpha1.KusciaDeploymentPartyStatus{
		Phase:               kusciav1alpha1.KusciaDeploymentPhaseProgressing,
		Role:                role,
		Replicas:            deployment.Status.Replicas,
		UpdatedReplicas:     deployment.Status.UpdatedReplicas,
		AvailableReplicas:   deployment.Status.AvailableReplicas,
		UnavailableReplicas: deployment.Status.UnavailableReplicas,
		Conditions:          deployment.Status.Conditions,
		CreationTimestamp:   &deployment.CreationTimestamp,
	}

	if curDepStatus.AvailableReplicas >= curDepStatus.Replicas {
		curDepStatus.Phase = kusciav1alpha1.KusciaDeploymentPhaseAvailable
	} else if curDepStatus.AvailableReplicas > 0 {
		curDepStatus.Phase = kusciav1alpha1.KusciaDeploymentPhasePartialAvailable
	}

	partyDepStatuses, ok := partyDeploymentStatuses[deployment.Namespace]
	if !ok {
		partyDepStatus := map[string]*kusciav1alpha1.KusciaDeploymentPartyStatus{
			deployment.Name: curDepStatus,
		}
		partyDeploymentStatuses[deployment.Namespace] = partyDepStatus
		return true
	}

	depStatus, ok := partyDepStatuses[deployment.Name]
	if !ok {
		partyDepStatuses[deployment.Name] = curDepStatus
		return true
	}

	if !reflect.DeepEqual(depStatus, curDepStatus) {
		partyDepStatuses[deployment.Name] = curDepStatus
		nlog.Infof("Party deployment %v/%v status changed from %+v to %+v", deployment.Namespace, deployment.Name, depStatus, curDepStatus)
		return true
	}
	return false
}

func (c *Controller) syncResources(ctx context.Context, partyKitInfos map[string]*PartyKitInfo) (err error) {
	if err = c.syncService(ctx, partyKitInfos); err != nil {
		return err
	}

	if err = c.syncConfigMap(ctx, partyKitInfos); err != nil {
		return err
	}

	if err = c.syncDeployment(ctx, partyKitInfos); err != nil {
		return err
	}

	return nil
}

func (c *Controller) syncService(ctx context.Context, partyKitInfos map[string]*PartyKitInfo) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error syncing service, %v", err)
		}
	}()

	for _, partyKitInfo := range partyKitInfos {
		for portName, serviceName := range partyKitInfo.dkInfo.portService {
			if _, err = c.serviceLister.Services(partyKitInfo.domainID).Get(serviceName); err != nil {
				if k8serrors.IsNotFound(err) {
					if err = c.createService(ctx, partyKitInfo, portName, serviceName); err != nil {
						partyKitInfo.kd.Status.Phase = kusciav1alpha1.KusciaDeploymentPhaseFailed
						partyKitInfo.kd.Status.Reason = string(createServiceFailed)
						partyKitInfo.kd.Status.Message = err.Error()
						return err
					}
					continue
				}
				return err
			}
		}
	}
	return nil
}

func (c *Controller) createService(ctx context.Context, partyKitInfo *PartyKitInfo, portName, serviceName string) error {
	ctrPort, ok := partyKitInfo.dkInfo.ports[portName]
	if !ok {
		return fmt.Errorf("container port %q is not found in deployment %q", portName, partyKitInfo.dkInfo.deploymentName)
	}

	service, err := generateService(partyKitInfo, serviceName, ctrPort)
	if err != nil {
		return fmt.Errorf("failed to generate service %v/%v for deployment %v, %v", service.Namespace, service.Name, partyKitInfo.dkInfo.deploymentName, err)
	}

	if _, err = c.kubeClient.CoreV1().Services(service.Namespace).Create(ctx, service, metav1.CreateOptions{}); err != nil {
		if k8serrors.IsAlreadyExists(err) {
			nlog.Warnf("Service %v/%v is already exists, %v", service.Namespace, service.Name, err)
			return nil
		}
		return fmt.Errorf("failed to create service %v/%v for deployment %v, %v", service.Namespace, service.Name, partyKitInfo.dkInfo.deploymentName, err)
	}
	return nil
}

func generateService(partyKitInfo *PartyKitInfo, serviceName string, port kusciav1alpha1.ContainerPort) (*corev1.Service, error) {
	labels := map[string]string{
		common.LabelController:               kusciaDeploymentName,
		common.LabelPortScope:                string(port.Scope),
		common.LabelInitiator:                partyKitInfo.kd.Spec.Initiator,
		common.LabelKusciaDeploymentUID:      string(partyKitInfo.kd.UID),
		common.LabelKusciaDeploymentName:     partyKitInfo.kd.Name,
		common.LabelKubernetesDeploymentName: partyKitInfo.dkInfo.deploymentName,
	}

	if partyKitInfo.interConnProtocol != "" {
		labels[common.LabelInterConnProtocolType] = partyKitInfo.interConnProtocol
	}

	if port.Scope != kusciav1alpha1.ScopeDomain {
		labels[common.LabelLoadBalancer] = string(common.DomainRouteLoadBalancer)
	}

	annotations := map[string]string{
		common.ProtocolAnnotationKey:     string(port.Protocol),
		common.AccessDomainAnnotationKey: partyKitInfo.portAccessDomains[port.Name],
	}

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        serviceName,
			Namespace:   partyKitInfo.domainID,
			Labels:      labels,
			Annotations: annotations,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(partyKitInfo.kd, kusciav1alpha1.SchemeGroupVersion.WithKind(kusciaDeploymentName)),
			},
		},
		Spec: corev1.ServiceSpec{
			Type:      corev1.ServiceTypeClusterIP,
			ClusterIP: "None",
			Selector: map[string]string{
				common.LabelKubernetesDeploymentName: partyKitInfo.dkInfo.deploymentName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:     port.Name,
					Port:     port.Port,
					Protocol: corev1.ProtocolTCP,
					TargetPort: intstr.IntOrString{
						Type:   intstr.String,
						StrVal: port.Name,
					},
				},
			},
		},
	}

	return svc, nil
}

func (c *Controller) syncConfigMap(ctx context.Context, partyKitInfos map[string]*PartyKitInfo) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error syncing configmap, %v", err)
		}
	}()

	for _, partyKitInfo := range partyKitInfos {
		if len(partyKitInfo.configTemplates) > 0 && partyKitInfo.configTemplatesCMName != "" {
			if _, err = c.configMapLister.ConfigMaps(partyKitInfo.domainID).Get(partyKitInfo.configTemplatesCMName); err != nil {
				if k8serrors.IsNotFound(err) {
					if err = c.createConfigMap(ctx, partyKitInfo); err != nil {
						partyKitInfo.kd.Status.Phase = kusciav1alpha1.KusciaDeploymentPhaseFailed
						partyKitInfo.kd.Status.Reason = string(createConfigMapFailed)
						partyKitInfo.kd.Status.Message = err.Error()
						return err
					}
					continue
				}
				return err
			}
		}
	}
	return nil
}

func (c *Controller) createConfigMap(ctx context.Context, partyKitInfo *PartyKitInfo) error {
	cm := generateConfigMap(partyKitInfo)
	if _, err := c.kubeClient.CoreV1().ConfigMaps(partyKitInfo.domainID).Create(ctx, cm, metav1.CreateOptions{}); err != nil {
		if k8serrors.IsAlreadyExists(err) {
			nlog.Warnf("Configmap %v/%v is already exists, %v", cm.Namespace, cm.Name, err)
			return nil
		}
		return fmt.Errorf("failed to create configmap %v/%v for deployment %v, %v", cm.Namespace, cm.Name, partyKitInfo.dkInfo.deploymentName, err)
	}
	return nil
}

func generateConfigMap(partyKitInfo *PartyKitInfo) *corev1.ConfigMap {
	labels := map[string]string{
		common.LabelController:               kusciaDeploymentName,
		common.LabelInitiator:                partyKitInfo.kd.Spec.Initiator,
		common.LabelKusciaDeploymentUID:      string(partyKitInfo.kd.UID),
		common.LabelKusciaDeploymentName:     partyKitInfo.kd.Name,
		common.LabelKubernetesDeploymentName: partyKitInfo.dkInfo.deploymentName,
	}

	if partyKitInfo.interConnProtocol != "" {
		labels[common.LabelInterConnProtocolType] = partyKitInfo.interConnProtocol
	}

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      partyKitInfo.configTemplatesCMName,
			Namespace: partyKitInfo.domainID,
			Labels:    labels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(partyKitInfo.kd, kusciav1alpha1.SchemeGroupVersion.WithKind(kusciaDeploymentName)),
			},
		},
		Data: partyKitInfo.configTemplates,
	}
}

func (c *Controller) syncDeployment(ctx context.Context, partyKitInfos map[string]*PartyKitInfo) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error syncing deployment, %v", err)
		}
	}()

	for _, partyKitInfo := range partyKitInfos {
		if _, err = c.deploymentLister.Deployments(partyKitInfo.domainID).Get(partyKitInfo.dkInfo.deploymentName); err != nil {
			if k8serrors.IsNotFound(err) {
				if err = c.createDeployment(ctx, partyKitInfo); err != nil {
					partyKitInfo.kd.Status.Phase = kusciav1alpha1.KusciaDeploymentPhaseFailed
					partyKitInfo.kd.Status.Reason = string(createDeploymentFailed)
					partyKitInfo.kd.Status.Message = err.Error()
					return err
				}
				continue
			}
			return err
		}

		if err = c.updateDeployment(ctx, partyKitInfo); err != nil {
			return err
		}
	}
	return nil
}

func (c *Controller) createDeployment(ctx context.Context, partyKitInfo *PartyKitInfo) error {
	deployment, err := c.generateDeployment(partyKitInfo)
	if err != nil {
		return err
	}

	if _, err = c.kubeClient.AppsV1().Deployments(partyKitInfo.domainID).Create(ctx, deployment, metav1.CreateOptions{}); err != nil {
		if k8serrors.IsAlreadyExists(err) {
			nlog.Warnf("Deployment %v is already exists, %v", partyKitInfo.dkInfo.deploymentName, err)
			return nil
		}
		return fmt.Errorf("failed to create deployment %v/%v for deployment %v, %v", deployment.Namespace, deployment.Name, partyKitInfo.dkInfo.deploymentName, err)
	}
	return err
}

func (c *Controller) generateDeployment(partyKitInfo *PartyKitInfo) (*appsv1.Deployment, error) {
	selectorLabels := map[string]string{
		common.LabelController:               kusciaDeploymentName,
		common.LabelInitiator:                partyKitInfo.kd.Spec.Initiator,
		common.LabelKusciaDeploymentUID:      string(partyKitInfo.kd.UID),
		common.LabelKusciaDeploymentName:     partyKitInfo.kd.Name,
		common.LabelKubernetesDeploymentName: partyKitInfo.dkInfo.deploymentName,
	}

	ns, err := c.namespaceLister.Get(partyKitInfo.domainID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate deployment %v, %v", partyKitInfo.dkInfo.deploymentName, err)
	}

	schedulerName := common.KusciaSchedulerName
	if ns.Labels != nil {
		if ns.Labels[common.LabelDomainRole] == string(kusciav1alpha1.Partner) {
			schedulerName = fmt.Sprintf("%v-%v", partyKitInfo.domainID, schedulerName)
		}
	}

	if partyKitInfo.interConnProtocol != "" {
		selectorLabels[common.LabelInterConnProtocolType] = partyKitInfo.interConnProtocol
	}

	maxSurge := intstr.FromString("25%")
	maxUnavailable := intstr.FromString("25%")
	updateStrategy := &appsv1.DeploymentStrategy{
		Type: appsv1.RollingUpdateDeploymentStrategyType,
		RollingUpdate: &appsv1.RollingUpdateDeployment{
			MaxSurge:       &maxSurge,
			MaxUnavailable: &maxUnavailable,
		},
	}
	if partyKitInfo.deployTemplate.Strategy != nil {
		updateStrategy = partyKitInfo.deployTemplate.Strategy
	}

	automountServiceAccountToken := false
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      partyKitInfo.dkInfo.deploymentName,
			Namespace: partyKitInfo.domainID,
			Labels:    selectorLabels,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(partyKitInfo.kd, kusciav1alpha1.SchemeGroupVersion.WithKind(kusciaDeploymentName)),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: partyKitInfo.deployTemplate.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: selectorLabels,
			},
			Strategy: *updateStrategy,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: selectorLabels,
				},
				Spec: corev1.PodSpec{
					Tolerations: []corev1.Toleration{
						{
							Key:      common.KusciaTaintTolerationKey,
							Operator: corev1.TolerationOpExists,
							Effect:   corev1.TaintEffectNoSchedule,
						},
					},
					NodeSelector: map[string]string{
						common.LabelNodeNamespace: partyKitInfo.domainID,
					},
					SchedulerName:                schedulerName,
					AutomountServiceAccountToken: &automountServiceAccountToken,
				},
			},
		},
	}

	renderConfigTemplateVolume := false
	for _, ctr := range partyKitInfo.deployTemplate.Spec.Containers {
		if ctr.ImagePullPolicy == "" {
			ctr.ImagePullPolicy = corev1.PullIfNotPresent
		}

		resCtr := corev1.Container{
			Name:                     ctr.Name,
			Image:                    partyKitInfo.dkInfo.image,
			Command:                  ctr.Command,
			Args:                     ctr.Args,
			WorkingDir:               ctr.WorkingDir,
			Env:                      ctr.Env,
			EnvFrom:                  ctr.EnvFrom,
			Resources:                ctr.Resources,
			LivenessProbe:            ctr.LivenessProbe,
			ReadinessProbe:           ctr.ReadinessProbe,
			StartupProbe:             ctr.StartupProbe,
			ImagePullPolicy:          ctr.ImagePullPolicy,
			SecurityContext:          ctr.SecurityContext,
			TerminationMessagePolicy: corev1.TerminationMessageFallbackToLogsOnError,
		}

		for _, port := range ctr.Ports {
			resCtr.Ports = append(resCtr.Ports, corev1.ContainerPort{
				Name:          port.Name,
				ContainerPort: port.Port,
				Protocol:      corev1.ProtocolTCP,
			})
		}

		protoJSONOptions := protojson.MarshalOptions{EmitUnpopulated: true}
		clusterDefine, err := protoJSONOptions.Marshal(partyKitInfo.dkInfo.clusterDef)
		if err != nil {
			return nil, fmt.Errorf("failed to generate deployment %v, %v", partyKitInfo.dkInfo.deploymentName, err)
		}

		allocatedPorts, err := protoJSONOptions.Marshal(partyKitInfo.dkInfo.allocatedPorts)
		if err != nil {
			return nil, fmt.Errorf("failed to generate deployment %v, %v", partyKitInfo.dkInfo.deploymentName, err)
		}

		resCtr.Env = append(resCtr.Env, []corev1.EnvVar{
			{
				Name:  common.EnvClusterDefine,
				Value: string(clusterDefine),
			},
			{
				Name:  common.EnvAllocatedPorts,
				Value: string(allocatedPorts),
			},
			{
				Name:  common.EnvInputConfig,
				Value: partyKitInfo.kd.Spec.InputConfig,
			},
		}...)

		if partyKitInfo.kd.Labels != nil && partyKitInfo.kd.Labels[common.LabelKusciaDeploymentAppType] == string(common.ServingAppType) {
			resCtr.Env = append(resCtr.Env, corev1.EnvVar{
				Name:  common.EnvServingID,
				Value: partyKitInfo.kd.Name,
			})
		}

		if len(ctr.ConfigVolumeMounts) > 0 && partyKitInfo.configTemplatesCMName != "" {
			renderConfigTemplateVolume = true
			for _, vm := range ctr.ConfigVolumeMounts {
				resCtr.VolumeMounts = append(resCtr.VolumeMounts, corev1.VolumeMount{
					Name:      configTemplateVolumeName,
					MountPath: vm.MountPath,
					SubPath:   vm.SubPath,
				})
			}
		}

		deployment.Spec.Template.Spec.Containers = append(deployment.Spec.Template.Spec.Containers, resCtr)
	}

	if renderConfigTemplateVolume {
		if deployment.Spec.Template.Annotations == nil {
			deployment.Spec.Template.Annotations = make(map[string]string)
		}
		deployment.Spec.Template.Annotations[common.ConfigTemplateVolumesAnnotationKey] = configTemplateVolumeName
		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, corev1.Volume{
			Name: configTemplateVolumeName,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: partyKitInfo.configTemplatesCMName,
					},
				},
			},
		})
	}
	return deployment, nil
}

func (c *Controller) updateDeployment(ctx context.Context, partyKitInfo *PartyKitInfo) error {
	deployment, err := c.deploymentLister.Deployments(partyKitInfo.domainID).Get(partyKitInfo.dkInfo.deploymentName)
	if err != nil {
		return fmt.Errorf("failed to check if need update deployment %v, %v", partyKitInfo.dkInfo.deploymentName, err)
	}

	deploymentCopy := deployment.DeepCopy()
	needUpdate := false
	for _, kdParty := range partyKitInfo.kd.Spec.Parties {
		if kdParty.DomainID == partyKitInfo.domainID && kdParty.Role == partyKitInfo.role {
			// check replicas
			if kdParty.Template.Replicas != nil && deploymentCopy.Spec.Replicas != nil && *kdParty.Template.Replicas != *deploymentCopy.Spec.Replicas {
				nlog.Debugf("Deployment %v/%v replicas changed from %v to %v", deploymentCopy.Namespace, deploymentCopy.Name, *deploymentCopy.Spec.Replicas, *kdParty.Template.Replicas)
				needUpdate = true
				deploymentCopy.Spec.Replicas = kdParty.Template.Replicas
			}

			// check strategy
			if kdParty.Template.Strategy != nil && !reflect.DeepEqual(*kdParty.Template.Strategy, deploymentCopy.Spec.Strategy) {
				nlog.Debugf("Deployment %v/%v strategy changed from %v to %v", deploymentCopy.Namespace, deploymentCopy.Name, deploymentCopy.Spec.Strategy, *kdParty.Template.Strategy)
				needUpdate = true
				deploymentCopy.Spec.Strategy = *kdParty.Template.Strategy
			}

			// check container image
			for i, ctr := range deploymentCopy.Spec.Template.Spec.Containers {
				if ctr.Image != partyKitInfo.dkInfo.image {
					nlog.Debugf("Deployment %v/%v pod container %v image changed from %v to %v", deploymentCopy.Namespace, deploymentCopy.Name, ctr.Name, ctr.Image, partyKitInfo.dkInfo.image)
					needUpdate = true
					deploymentCopy.Spec.Template.Spec.Containers[i].Image = partyKitInfo.dkInfo.image
				}
			}

			// check container resources
			for _, pc := range partyKitInfo.deployTemplate.Spec.Containers {
				for i, dc := range deploymentCopy.Spec.Template.Spec.Containers {
					if pc.Name == dc.Name {
						if !reflect.DeepEqual(pc.Resources, dc.Resources) {
							nlog.Debugf("Deployment %v/%v pod container %v resources changed from %v to %v", deploymentCopy.Namespace, deploymentCopy.Name, pc.Name, dc.Resources.String(), pc.Resources.String())
							needUpdate = true
							deploymentCopy.Spec.Template.Spec.Containers[i].Resources = *pc.Resources.DeepCopy()
						}
					}
				}
			}

			// check input_config
			envExist := false
			for i, ctr := range deploymentCopy.Spec.Template.Spec.Containers {
				for j, ctrEnv := range ctr.Env {
					if ctrEnv.Name == common.EnvInputConfig {
						envExist = true
						if partyKitInfo.kd.Spec.InputConfig != ctrEnv.Value {
							nlog.Debugf("Deployment %v/%v pod container %v env %v changed from %v to %v", deploymentCopy.Namespace, deploymentCopy.Name, ctr.Name, common.EnvInputConfig, ctrEnv.Value, partyKitInfo.kd.Spec.InputConfig)
							needUpdate = true
							deploymentCopy.Spec.Template.Spec.Containers[i].Env[j].Value = partyKitInfo.kd.Spec.InputConfig
						}
					}
				}
			}
			if !envExist {
				nlog.Debugf("Deployment %v/%v pod containers need add env %v:%v", deploymentCopy.Namespace, deploymentCopy.Name, common.EnvInputConfig, partyKitInfo.kd.Spec.InputConfig)
				needUpdate = true
				for index := range deploymentCopy.Spec.Template.Spec.Containers {
					deploymentCopy.Spec.Template.Spec.Containers[index].Env = append(deploymentCopy.Spec.Template.Spec.Containers[index].Env, corev1.EnvVar{
						Name:  common.EnvInputConfig,
						Value: partyKitInfo.kd.Spec.InputConfig,
					})
				}
			}
		}
	}

	if needUpdate {
		_, err = c.kubeClient.AppsV1().Deployments(deploymentCopy.Namespace).Update(ctx, deploymentCopy, metav1.UpdateOptions{})
		if err != nil {
			return fmt.Errorf("failed to update deployment %v/%v, %v", deploymentCopy.Namespace, deployment.Name, err)
		}
	}

	return nil
}
