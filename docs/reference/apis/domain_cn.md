# Domain

在 Kuscia 中将隐私计算的节点称为 Domain，一个 Domain 中可以包含多个 K3s
的工作节点（Node）。详情请参考 [Domain](../concepts/domain_cn.md) 。
你可以从 [这里](https://github.com/secretflow/kuscia/tree/main/proto/api/v1alpha1/kusciaapi/domain.proto) 找到对应的
protobuf 文件。

## 接口总览

| 方法名                                               | 请求类型                      | 响应类型                       | 描述             |
| ---------------------------------------------------- | ----------------------------- | ------------------------------ | ---------------- |
| [CreateDomain](#create-domain)                       | CreateDomainRequest           | CreateDomainResponse           | 创建节点         |
| [UpdateDomain](#update-domain)                       | UpdateDomainRequest           | UpdateDomainResponse           | 更新节点         |
| [DeleteDomain](#delete-domain)                       | DeleteDomainRequest           | DeleteDomainResponse           | 删除节点         |
| [QueryDomain](#query-domain)                         | QueryDomainRequest            | QueryDomainResponse            | 查询节点         |
| [BatchQueryDomainStatus](#batch-query-domain-status) | BatchQueryDomainStatusRequest | BatchQueryDomainStatusResponse | 批量查询节点状态 |

## 接口详情

{#create-domain}

### 创建节点

#### HTTP 路径

/api/v1/domain/create

#### 请求（CreateDomainRequest）

| 字段        | 类型                                         | 可选 | 描述                                                                 |
| ----------- | -------------------------------------------- | ---- | -------------------------------------------------------------------- |
| header      | [RequestHeader](summary_cn.md#requestheader) | 是   | 自定义请求内容                                                       |
| domain_id   | string                                       | 否   | 节点 ID 需要符合 DNS 子域名规则要求，参考 [DomainId 规则要求](https://kubernetes.io/zh-cn/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names)                                                             |
| role        | string                                       | 是   | 角色：\["", "partner"]，参考 [Domain 概念](../concepts/domain_cn.md) |
| cert        | string                                       | 是   | BASE64 的计算节点证书，参考 [Domain 概念](../concepts/domain_cn.md)  |
| auth_center | [AuthCenter](#auth-center)                   | 是   | 节点到中心的授权模式                                                 |

#### 响应（CreateDomainResponse）

| 字段   | 类型                           | 可选 | 描述     |
| ------ | ------------------------------ | ---- | -------- |
| status | [Status](summary_cn.md#status) | 否   | 状态信息 |

{#update-domain}

### 更新节点

#### HTTP 路径

/api/v1/domain/update

#### 请求（UpdateDomainRequest）

| 字段      | 类型                                         | 可选 | 描述                                                                 |
| --------- | -------------------------------------------- | ---- | -------------------------------------------------------------------- |
| header    | [RequestHeader](summary_cn.md#requestheader) | 是   | 自定义请求内容                                                       |
| domain_id | string                                       | 是   | 节点 ID                                                              |
| role      | string                                       | 是   | 角色：\["", "partner"]，参考 [Domain 概念](../concepts/domain_cn.md) |
| cert      | string                                       | 是   | BASE64 的计算节点证书，参考 [Domain 概念](../concepts/domain_cn.md)  |

#### 响应（UpdateDomainResponse）

| 字段   | 类型                           | 可选 | 描述     |
| ------ | ------------------------------ | ---- | -------- |
| status | [Status](summary_cn.md#status) | 否   | 状态信息 |

{#delete-domain}

### 删除节点

#### HTTP 路径

/api/v1/domain/delete

#### 请求（DeleteDomainRequest）

| 字段      | 类型                                         | 可选 | 描述           |
| --------- | -------------------------------------------- | ---- | -------------- |
| header    | [RequestHeader](summary_cn.md#requestheader) | 是   | 自定义请求内容 |
| domain_id | string                                       | 否   | 节点 ID        |

#### 响应（DeleteDomainResponse）

| 字段   | 类型                           | 可选 | 描述     |
| ------ | ------------------------------ | ---- | -------- |
| status | [Status](summary_cn.md#status) | 否   | 状态信息 |

{#query-domain}

### 查询节点

#### HTTP 路径

/api/v1/domain/query

#### 请求（QueryDomainRequest）

| 字段      | 类型                                         | 可选 | 描述           |
| --------- | -------------------------------------------- | ---- | -------------- |
| header    | [RequestHeader](summary_cn.md#requestheader) | 是   | 自定义请求内容 |
| domain_id | string                                       | 否   | 节点 ID        |

#### 响应（QueryDomainResponse）

| 字段                       | 类型                                        | 可选 | 描述                                                                 |
| -------------------------- | ------------------------------------------- | ---- | -------------------------------------------------------------------- |
| status                     | [Status](summary_cn.md#status)              | 否   | 状态信息                                                             |
| data                       | QueryDomainResponseData                     |      |                                                                      |
| data.domain_id             | string                                      | 否   | 节点 ID                                                              |
| data.role                  | string                                      | 否   | 角色：\["", "partner"]，参考 [Domain 概念](../concepts/domain_cn.md) |
| data.cert                  | string                                      | 是   | BASE64 的计算节点证书，参考 [Domain 概念](../concepts/domain_cn.md)  |
| data.annotations           | map[string]string                           | 是   | 节点的额外信息，比如是否是内置节点等                                 |
| data.auth_center           | [AuthCenter](#auth-center)                  | 是   | 节点到中心的授权模式                                                 |
| data.node_statuses         | [NodeStatus](#node-status)[]                | 否   | 物理节点状态                                                         |
| data.deploy_token_statuses | [DeployTokenStatus](#deploy-token-status)[] | 否   | 部署令牌状态                                                         |

{#batch-query-domain-status}

### 批量查询节点状态

#### HTTP 路径

/api/v1/domain/status/batchQuery

#### 请求（BatchQueryDomainStatusRequest）

| 字段       | 类型                                         | 可选 | 描述                 |
| ---------- | -------------------------------------------- | ---- | -------------------- |
| header     | [RequestHeader](summary_cn.md#requestheader) | 是   | 自定义请求内容       |
| domain_ids | string[]                                     | 否   | 待查询的节点 ID 列表 |

#### 响应（ BatchQueryDomainStatusResponse）

| 字段         | 类型                               | 可选 | 描述     |
| ------------ | ---------------------------------- | ---- | -------- |
| status       | [Status](summary_cn.md#status)     | 否   | 状态信息 |
| data         | BatchQueryDomainStatusResponseData | 否   |          |
| data.domains | [DomainStatus](#domain-status)[]   | 否   | 节点列表 |

## 公共

{#domain-status}

### DomainStatus

| 字段                  | 类型                                        | 可选 | 描述             |
| --------------------- | ------------------------------------------- | ---- | ---------------- |
| domain_id             | string                                      | 否   | 节点 ID          |
| node_statuses         | [NodeStatus](#node-status)[]                | 否   | 真实物理节点状态 |
| deploy_token_statuses | [DeployTokenStatus](#deploy-token-status)[] | 否   | 部署令牌状态     |

{#node-status}

### NodeStatus

| 字段                 | 类型   | 可选 | 描述                                                    |
| -------------------- | ------ | ---- | ------------------------------------------------------- |
| name                 | string | 否   | 节点名称                                                |
| status               | string | 否   | 节点状态                                                |
| version              | string | 否   | 节点 Agent 版本                                         |
| last_heartbeat_time  | string | 否   | 最后心跳时间，RFC3339 格式（e.g. 2006-01-02T15:04:05Z） |
| last_transition_time | string | 否   | 最后更新时间，RFC3339 格式（e.g. 2006-01-02T15:04:05Z） |

{#auth-center}

### AuthCenter

| 字段                | 类型   | 可选 | 描述                                              |
| ------------------- | ------ | ---- | ------------------------------------------------- |
| authentication_type | string | 否   | 节点到中心授权类型，目前仅支持 Token              |
| token_gen_method    | string | 否   | 节点到中心 Token 生成类型，目前仅支持 UID-RSA-GEN |

{#deploy-token-status}

### DeployTokenStatus

| 字段                 | 类型   | 可选 | 描述                                                    |
| -------------------- | ------ | ---- | ------------------------------------------------------- |
| token                | string | 否   | 部署令牌                                                |
| state                | string | 否   | 部署令牌状态 used, unsed                                |
| last_transition_time | string | 否   | 最后更新时间，RFC3339 格式（e.g. 2006-01-02T15:04:05Z） |
