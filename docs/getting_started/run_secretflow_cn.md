# Kuscia 入门教程 —— 了解 KusciaJob

你将会尝试自己配置并且提交一个作业，在这个过程中，你将会认识 Kuscia 中的重要概念 —— **KusciaJob**。

## 准备集群

这里假设你已经在机器上部署并且启动了一个示例集群。如果你还没有这样做，请先按照 [快速开始][quickstart] 的指引启动一个集群。

[quickstart]: ./quickstart_cn.md

## 准备数据

你可以使用 Kuscia 中自带的数据文件，或者使用你自己的数据文件。

在 Kuscia 中，节点数据文件的存放路径为节点容器的`/home/kuscia/var/storage`，你可以在容器中查看这个数据文件。

### 查看 Kuscia 示例数据

这里以 alice 节点为例，首先进入节点容器：

```shell
docker exec -it ${USER}-kuscia-lite-alice bash
```

在 alice 节点容器中查看节点示例数据：

```shell
cat /home/kuscia/var/storage/data/alice.csv
```

bob 节点同理。

{#prepare-your-own-data}

### 准备你自己的数据

你也可以使用你自己的数据文件，首先你要你的数据文件复制到节点容器中，还是以 alice 节点为例：

```shell
docker cp {your_alice_data} ${USER}-kuscia-lite-alice:/home/kuscia/var/storage/data/
```

接下来你可以像 [查看 Kuscia 示例数据](#kuscia) 一样查看你的数据文件，这里不再赘述。

{#configure-kuscia-job}

## 配置 KusciaJob

我们需要在 kuscia-master 节点容器中配置和运行 Job，首先，让我们先进入 kuscia-master 节点容器：

```shell
docker exec -it ${USER}-kuscia-master bash
```

### 使用 Kuscia 示例数据配置 KusciaJob

下面的示例展示了一个 KusciaJob， 该任务流完成 2 个任务：

1. job-psi 读取 alice 和 bob 的数据文件，进行隐私求交，求交的结果分别保存为两个参与方的`psi-output.csv`。
2. job-split 读取 alice 和 bob 上一步中求交的结果文件，并拆分成训练集和测试集，分别保存为两个参与方的`train-dataset.csv`、`test-dataset.csv`。

这个 KusciaJob 的名称为 job-best-effort-linear，在一个 Kuscia 集群中，这个名称必须是唯一的，由`.metadata.name`指定。

在 kuscia-master 容器中，在任意路径创建文件 job-best-effort-linear.yaml，内容如下：

```yaml
apiVersion: kuscia.secretflow/v1alpha1
kind: KusciaJob
metadata:
  name: job-best-effort-linear
spec:
  initiator: alice
  scheduleMode: BestEffort
  maxParallelism: 2
  tasks:
    - taskID: job-psi
      alias: job-psi
      priority: 100
      taskInputConfig: '{"sf_datasource_config":{"bob":{"id":"default-data-source"},"alice":{"id":"default-data-source"}},"sf_cluster_desc":{"parties":["alice","bob"],"devices":[{"name":"spu","type":"spu","parties":["alice","bob"],"config":"{\"runtime_config\":{\"protocol\":\"REF2K\",\"field\":\"FM64\"},\"link_desc\":{\"connect_retry_times\":60,\"connect_retry_interval_ms\":1000,\"brpc_channel_protocol\":\"http\",\"brpc_channel_connection_type\":\"pooled\",\"recv_timeout_ms\":1200000,\"http_timeout_ms\":1200000}}"},{"name":"heu","type":"heu","parties":["alice","bob"],"config":"{\"mode\": \"PHEU\", \"schema\": \"paillier\", \"key_size\": 2048}"}],"ray_fed_config":{"cross_silo_comm_backend":"brpc_link"}},"sf_node_eval_param":{"domain":"preprocessing","name":"psi","version":"0.0.1","attr_paths":["input/receiver_input/key","input/sender_input/key","protocol","precheck_input","bucket_size","curve_type"],"attrs":[{"ss":["id1"]},{"ss":["id2"]},{"s":"ECDH_PSI_2PC"},{"b":true},{"i64":"1048576"},{"s":"CURVE_FOURQ"}],"inputs":[{"type":"sf.table.individual","meta":{"@type":"type.googleapis.com/secretflow.component.IndividualTable","schema":{"ids":["id1"],"features":["age","education","default","balance","housing","loan","day","duration","campaign","pdays","previous","job_blue-collar","job_entrepreneur","job_housemaid","job_management","job_retired","job_self-employed","job_services","job_student","job_technician","job_unemployed","marital_divorced","marital_married","marital_single"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float"]},"num_lines":"-1"},"data_refs":[{"uri":"alice.csv","party":"alice","format":"csv"}]},{"type":"sf.table.individual","meta":{"@type":"type.googleapis.com/secretflow.component.IndividualTable","schema":{"ids":["id2"],"features":["contact_cellular","contact_telephone","contact_unknown","month_apr","month_aug","month_dec","month_feb","month_jan","month_jul","month_jun","month_mar","month_may","month_nov","month_oct","month_sep","poutcome_failure","poutcome_other","poutcome_success","poutcome_unknown"],"labels":["y"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float"],"label_types":["int"]},"num_lines":"-1"},"data_refs":[{"uri":"bob.csv","party":"bob","format":"csv"}]}]},"sf_output_uris":["psi-output.csv"],"sf_output_ids":["psi-output"]}'
      appImage: secretflow-image
      parties:
        - domainID: alice
        - domainID: bob
    - taskID: job-split
      alias: job-split
      priority: 100
      dependencies: ['job-psi']
      taskInputConfig: '{"sf_datasource_config":{"bob":{"id":"default-data-source"},"alice":{"id":"default-data-source"}},"sf_cluster_desc":{"parties":["alice","bob"],"devices":[{"name":"spu","type":"spu","parties":["alice","bob"],"config":"{\"runtime_config\":{\"protocol\":\"REF2K\",\"field\":\"FM64\"},\"link_desc\":{\"connect_retry_times\":60,\"connect_retry_interval_ms\":1000,\"brpc_channel_protocol\":\"http\",\"brpc_channel_connection_type\":\"pooled\",\"recv_timeout_ms\":1200000,\"http_timeout_ms\":1200000}}"},{"name":"heu","type":"heu","parties":["alice","bob"],"config":"{\"mode\": \"PHEU\", \"schema\": \"paillier\", \"key_size\": 2048}"}],"ray_fed_config":{"cross_silo_comm_backend":"brpc_link"}},"sf_node_eval_param":{"domain":"preprocessing","name":"train_test_split","version":"0.0.1","attr_paths":["train_size","test_size","random_state","shuffle"],"attrs":[{"f":0.75},{"f":0.25},{"i64":1234},{"b":true}],"inputs":[{"type":"sf.table.vertical_table","meta":{"@type":"type.googleapis.com/secretflow.component.VerticalTable","schemas":[{"ids":["id1"],"features":["age","education","default","balance","housing","loan","day","duration","campaign","pdays","previous","job_blue-collar","job_entrepreneur","job_housemaid","job_management","job_retired","job_self-employed","job_services","job_student","job_technician","job_unemployed","marital_divorced","marital_married","marital_single"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float"]},{"ids":["id2"],"features":["contact_cellular","contact_telephone","contact_unknown","month_apr","month_aug","month_dec","month_feb","month_jan","month_jul","month_jun","month_mar","month_may","month_nov","month_oct","month_sep","poutcome_failure","poutcome_other","poutcome_success","poutcome_unknown","y"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","int"]}]},"data_refs":[{"uri":"psi-output.csv","party":"alice","format":"csv"},{"uri":"psi-output.csv","party":"bob","format":"csv"}]}]},"sf_output_uris":["train-dataset.csv","test-dataset.csv"],"sf_output_ids":["train-dataset","test-dataset"]}'
      appImage: secretflow-image
      parties:
        - domainID: alice
        - domainID: bob
```

:::{tip}

更多有关 KusciaJob 配置的信息，请查看 [KusciaJob](../reference/concepts/kusciajob_cn.md)。

KusciaJob 的算子参数由 `taskInputConfig` 字段定义，对于不同的算子，算子的参数不同。

本教程使用的是 SecretFlow 的算子参数定义。关于该定义的更多细节，敬请查阅 [SecretFlow 的对应文档][secretflow-component-spec]。

[secretflow-component-spec]: https://www.secretflow.org.cn/docs/secretflow/latest/zh-Hans/component

:::

### 使用你自己的数据配置 KusciaJob

如果你要使用你自己的数据，可以将两个算子中的`taskInputConfig.sf_node_eval_param`中的`inputs`和`output_uris`中的数据文件路径修改为你在 [准备你自己的数据](#prepare-your-own-data) 中的数据文件目标路径即可。

## 运行 KusciaJob

现在我们已经配置好了一个 KusciaJob ，接下来，让我们运行这个 KusciaJob， 在 kuscia-master 容器中执行 ：

```shell
kubectl apply -f job-best-effort-linear.yaml
```

## 查看 KusciaJob 运行状态

现在我们提交了这个 KusciaJob ，接下来我们可以在 kuscia-master 容器中通过下面的命令查看 KusciaJob 的运行情况。

### 查看所有的 KusciaJob

```shell
kubectl get kj
```

你可以看到如下输出：

```shell
NAME                     STARTTIME   COMPLETIONTIME   LASTRECONCILETIME   PHASE
job-best-effort-linear   3s                           3s                  Running
```

job-best-effort-linear 就是我们刚刚创建出来的 KusciaJob 。

### 查看运行中的某个 KusciaJob 的详细状态

通过指定`-o yaml`参数，我们可以以 Yaml 的形式看到 KusciaJob 的详细状态。job-best-effort-linear 是你在 [配置 Job](#configure-kuscia-job) 中指定的 KusciaJob 的名称。

```shell
kubectl get kj job-best-effort-linear -o yaml
```

如果任务成功了，你可以看到如下输出：

```shell
apiVersion: kuscia.secretflow/v1alpha1
kind: KusciaJob
metadata:
  creationTimestamp: "2023-03-30T12:11:41Z"
  generation: 1
  name: job-best-effort-linear
  resourceVersion: "19002"
  uid: 085e10e6-5d3e-43cf-adaa-715d76a6af9b
spec:
  initiator: alice
  maxParallelism: 2
  scheduleMode: BestEffort
  tasks:
  - appImage: secretflow-image
    parties:
    - domainID: alice
    - domainID: bob
    priority: 100
    taskID: job-psi
    alias: job-psi
    taskInputConfig: '{"sf_datasource_config":{"bob":{"id":"default-data-source"},"alice":{"id":"default-data-source"}},"sf_cluster_desc":{"parties":["alice","bob"],"devices":[{"name":"spu","type":"spu","parties":["alice","bob"],"config":"{\"runtime_config\":{\"protocol\":\"REF2K\",\"field\":\"FM64\"},\"link_desc\":{\"connect_retry_times\":60,\"connect_retry_interval_ms\":1000,\"brpc_channel_protocol\":\"http\",\"brpc_channel_connection_type\":\"pooled\",\"recv_timeout_ms\":1200000,\"http_timeout_ms\":1200000}}"},{"name":"heu","type":"heu","parties":["alice","bob"],"config":"{\"mode\": \"PHEU\", \"schema\": \"paillier\", \"key_size\": 2048}"}],"ray_fed_config":{"cross_silo_comm_backend":"brpc_link"}},"sf_node_eval_param":{"domain":"preprocessing","name":"psi","version":"0.0.1","attr_paths":["input/receiver_input/key","input/sender_input/key","protocol","precheck_input","bucket_size","curve_type"],"attrs":[{"ss":["id1"]},{"ss":["id2"]},{"s":"ECDH_PSI_2PC"},{"b":true},{"i64":"1048576"},{"s":"CURVE_FOURQ"}],"inputs":[{"type":"sf.table.individual","meta":{"@type":"type.googleapis.com/secretflow.component.IndividualTable","schema":{"ids":["id1"],"features":["age","education","default","balance","housing","loan","day","duration","campaign","pdays","previous","job_blue-collar","job_entrepreneur","job_housemaid","job_management","job_retired","job_self-employed","job_services","job_student","job_technician","job_unemployed","marital_divorced","marital_married","marital_single"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float"]},"num_lines":"-1"},"data_refs":[{"uri":"alice.csv","party":"alice","format":"csv"}]},{"type":"sf.table.individual","meta":{"@type":"type.googleapis.com/secretflow.component.IndividualTable","schema":{"ids":["id2"],"features":["contact_cellular","contact_telephone","contact_unknown","month_apr","month_aug","month_dec","month_feb","month_jan","month_jul","month_jun","month_mar","month_may","month_nov","month_oct","month_sep","poutcome_failure","poutcome_other","poutcome_success","poutcome_unknown"],"labels":["y"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float"],"label_types":["int"]},"num_lines":"-1"},"data_refs":[{"uri":"bob.csv","party":"bob","format":"csv"}]}]},"sf_output_uris":["psi-output.csv"],"sf_output_ids":["psi-output"]}'
    tolerable: false
  - appImage: secretflow-image
    dependencies:
    - job-psi
    parties:
    - domainID: alice
    - domainID: bob
    priority: 100
    taskID: job-split
    alias: job-split
    taskInputConfig: '{"sf_datasource_config":{"bob":{"id":"default-data-source"},"alice":{"id":"default-data-source"}},"sf_cluster_desc":{"parties":["alice","bob"],"devices":[{"name":"spu","type":"spu","parties":["alice","bob"],"config":"{\"runtime_config\":{\"protocol\":\"REF2K\",\"field\":\"FM64\"},\"link_desc\":{\"connect_retry_times\":60,\"connect_retry_interval_ms\":1000,\"brpc_channel_protocol\":\"http\",\"brpc_channel_connection_type\":\"pooled\",\"recv_timeout_ms\":1200000,\"http_timeout_ms\":1200000}}"},{"name":"heu","type":"heu","parties":["alice","bob"],"config":"{\"mode\": \"PHEU\", \"schema\": \"paillier\", \"key_size\": 2048}"}],"ray_fed_config":{"cross_silo_comm_backend":"brpc_link"}},"sf_node_eval_param":{"domain":"preprocessing","name":"train_test_split","version":"0.0.1","attr_paths":["train_size","test_size","random_state","shuffle"],"attrs":[{"f":0.75},{"f":0.25},{"i64":1234},{"b":true}],"inputs":[{"type":"sf.table.vertical_table","meta":{"@type":"type.googleapis.com/secretflow.component.VerticalTable","schemas":[{"ids":["id1"],"features":["age","education","default","balance","housing","loan","day","duration","campaign","pdays","previous","job_blue-collar","job_entrepreneur","job_housemaid","job_management","job_retired","job_self-employed","job_services","job_student","job_technician","job_unemployed","marital_divorced","marital_married","marital_single"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float"]},{"ids":["id2"],"features":["contact_cellular","contact_telephone","contact_unknown","month_apr","month_aug","month_dec","month_feb","month_jan","month_jul","month_jun","month_mar","month_may","month_nov","month_oct","month_sep","poutcome_failure","poutcome_other","poutcome_success","poutcome_unknown","y"],"id_types":["str"],"feature_types":["float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","float","int"]}]},"data_refs":[{"uri":"psi-output.csv","party":"alice","format":"csv"},{"uri":"psi-output.csv","party":"bob","format":"csv"}]}]},"sf_output_uris":["train-dataset.csv","test-dataset.csv"],"sf_output_ids":["train-dataset","test-dataset"]}'
    tolerable: false
status:
  completionTime: "2023-03-30T12:12:07Z"
  conditions:
  - lastTransitionTime: "2023-03-30T12:11:41Z"
     status: "True"
    type: JobValidated
  - lastTransitionTime: "2023-03-30T12:11:41Z"
    status: "True"
    type: JobCreateInitialized
  - lastTransitionTime: "2023-03-30T12:11:42Z"
    status: "True"
    type: JobCreateSucceeded
  - lastTransitionTime: "2023-03-30T12:11:44Z"
    status: "True"
    type: JobStartInitialized
  - lastTransitionTime: "2023-03-30T12:11:44Z"
    status: "True"
    type: JobStartSucceeded
  lastReconcileTime: "2023-03-30T12:12:07Z"
  phase: Succeeded
  startTime: "2023-03-30T12:11:41Z"
  taskStatus:
    job-psi: Succeeded
    job-split: Succeeded
```

`status`字段记录了 KusciaJob 的运行状态，`.status.phase`字段描述了 KusciaJob 的整体状态，而 `.status.taskStatus`则描述了每个 KusciaTask 的状态。
详细信息请参考 [KusciaJob](../reference/concepts/kusciajob_cn.md) 。

### 查看 KusciaJob 中的某个 KusciaTask 的详细状态。

KusciaJob 中的每一个 KusciaTask 都有一个`taskID`，通过`taskID`我们可以查看某个 KusciaTask 的详细状态。

```shell
kubectl get kt job-psi -o yaml
```

KusciaTask 的信息这里不再赘述，请查看 [KusciaTask](../reference/concepts/kusciatask_cn.md) 。

## 删除 KusciaJob

当你想清理这个 KusciaJob 时，你可以通过下面的命令完成：

```shell
kubectl delete kj job-best-effort-linear
```

当这个 KusciaJob 被清理时， 这个 KusciaJob 创建的 KusciaTask 也会一起被清理。

## 接下来

恭喜！你已经完成了 Kuscia 的入门教程。接下来，你可以：

- 进一步阅读 [Kuscia 架构细节][architecture]，了解 Kuscia 的设计思路和概念。
- 了解 [Kuscia API][kuscia-api]。Kuscia API 是 Kuscia 的一个更上层封装，支持更方便地将 Kuscia 集成到其他系统中。
- 了解 [多机器部署][deploy-p2p] 的更多信息。
- 尝试运行其它算法或是引擎的作业，比如 [互联互通银联 BFIA 协议作业][tutorial-bfia] 或是 [FATE 作业][tutorial-fate]。

[architecture]: ../reference/architecture_cn.md
[kuscia-api]: ../reference/apis/summary_cn.md
[deploy-p2p]: ../deployment/deploy_p2p_cn.md
[tutorial-bfia]: ../tutorial/run_bfia_job_cn.md
[tutorial-fate]: ../tutorial/run_fate_cn.md

:::{tip}

如果你想要停止并清理入门教程使用的示例集群，可以查阅 [相关指引][stop-and-uninstall]。

[stop-and-uninstall]: ./quickstart_cn.md#停止体验集群

:::
