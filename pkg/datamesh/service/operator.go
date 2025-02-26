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

package service

import (
	"context"
	"path"
	"time"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/secretflow/kuscia/pkg/common"
	"github.com/secretflow/kuscia/pkg/datamesh/config"
	"github.com/secretflow/kuscia/pkg/utils/nlog"
	"github.com/secretflow/kuscia/pkg/web/utils"
	"github.com/secretflow/kuscia/proto/api/v1alpha1/datamesh"
)

func GetDefaultDataSourceID() string {
	return common.DefaultDataSourceID
}

type IOperatorService interface {
	Start(ctx context.Context)
}

type operatorService struct {
	conf          *config.DataMeshConfig
	datasourceSvc IDomainDataSourceService
}

func NewOperatorService(config *config.DataMeshConfig) IOperatorService {
	return &operatorService{
		conf:          config,
		datasourceSvc: NewDomainDataSourceService(config),
	}
}

func (o *operatorService) Start(ctx context.Context) {
	nlog.Infof("DataMesh operator service start")
	// register datasource with best efforts
	o.registerDatasourceBestEfforts()
	// do other logic
	go func() {
		o.doLogic(ctx)
	}()
}

func (o *operatorService) doLogic(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			nlog.Infof("DataMesh operator service stop")
			return
			//default:
			// TODO： check the status of domain data periodically
		}
	}
}

func (o *operatorService) registerDatasourceBestEfforts() {
	go func() {
		for !o.registerDefaultDatasource() {
			time.Sleep(5 * time.Second)
		}
	}()
}

func (o *operatorService) registerDefaultDatasource() bool {
	datasource, err := o.conf.KusciaClient.KusciaV1alpha1().DomainDataSources(o.conf.KubeNamespace).Get(context.Background(), GetDefaultDataSourceID(), metav1.GetOptions{})
	if err != nil {
		if k8serrors.IsNotFound(err) {
			req := &datamesh.CreateDomainDataSourceRequest{
				DatasourceId: GetDefaultDataSourceID(),
				Name:         GetDefaultDataSourceID(),
				Type:         common.DomainDataSourceTypeLocalFS,
				Info: &datamesh.DataSourceInfo{
					Localfs: &datamesh.LocalDataSourceInfo{
						Path: path.Join(o.conf.RootDir, common.DefaultDomainDataSourceLocalFSPath),
					},
					Oss: nil,
				},
			}
			nlog.Infof("Create default datasource.")
			resp := o.datasourceSvc.CreateDomainDataSource(context.Background(), req)
			nlog.Debugf("Create default datasource response code:%d , message:%s.", resp.Status.Code, resp.Status.Message)
			if resp.Status.Code != utils.ResponseCodeSuccess {
				nlog.Errorf("Create default datasource failed,code %d,error message:%s.", resp.Status.Code, resp.Status.Message)
			}
		} else {
			nlog.Errorf("Get default datasource failed, error:%s.", err.Error())
		}
	}
	if datasource != nil && datasource.Name == GetDefaultDataSourceID() {
		nlog.Infof("Datasource has been created successful.")
		return true
	}
	return false
}
