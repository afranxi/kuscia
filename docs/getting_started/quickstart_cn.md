# Kuscia 入门教程 —— 快速开始

你将会在单台机器上准备 Kuscia 需要的环境、快速部署一个示例 Kuscia 集群，然后尝试运行一个 [SecretFlow] 作业。

[SecretFlow]: https://www.secretflow.org.cn/docs/secretflow

## 部署模式：说明

在部署时有两种组网模式可供选择：

- [中心化组网模式](../reference/architecture_cn.md#中心化组网模式)：启动一个控制平面（master）容器和两个 Lite 节点（alice 和 bob）容器
- [点对点组网模式](../reference/architecture_cn.md#点对点组网模式)：启动两个 Autonomy 节点（alice 和 bob）容器

你可以选择其中任意一种或两种模式进行部署体验，在单台机器上可以同时部署两种模式。

## 环境

### 机器

操作系统：macOS, CentOS7, CentOS8, Ubuntu 16.04 及以上版本, Windows(Ubuntu on WSL2)

资源：8 core / 16G memory / 200G hard disk

CPU 架构：x86

### 环境准备

Kuscia 的部署需要依赖 docker 环境，docker 的安装请参考[官方文档](https://docs.docker.com/engine/install/)。以下为 CentOS 系统安装 docker 的示例：

```bash
# 安装 docker。
yum install -y yum-utils
yum-config-manager \
	--add-repo \
	https://download.docker.com/linux/centos/docker-ce.repo
yum install -y docker-ce docker-ce-cli containerd.io

# 启动 docker。
systemctl start docker
```

### 关于 macOS

macOS 默认给单个 docker container 分配了 2G 内存，请参考[官方文档](https://docs.docker.com/desktop/settings/mac/)将内存上限提高为 6G（Kuscia 2G + SecretFlow 4G) 。

此外，Kuscia 当前不支持 M1/M2 芯片的 Mac。

## 部署体验

### 前置操作

配置 Kuscia 镜像，以下示例选择使用 latest 版本镜像（更多镜像版本请参考 [Kuscia tags](https://hub.docker.com/r/secretflow/kuscia/tags)）：

```bash
# Docker Hub 镜像
export KUSCIA_IMAGE=secretflow/kuscia

# 阿里云镜像（推荐国内用户使用）
export KUSCIA_IMAGE=secretflow-registry.cn-hangzhou.cr.aliyuncs.com/secretflow/kuscia
```

获取 Kuscia 安装脚本，安装脚本会下载到当前目录：

```
docker run --rm --pull always $KUSCIA_IMAGE cat /home/kuscia/scripts/deploy/start_standalone.sh > start_standalone.sh && chmod u+x start_standalone.sh
```

### 中心化组网模式

```bash
# 启动集群，会拉起 3 个 docker 容器，包括一个控制平面 master 和两个 Lite 节点 alice 和 bob。
./start_standalone.sh center

# 登入 master 容器。
docker exec -it ${USER}-kuscia-master bash

# 创建并启动作业（两方 PSI 任务）。
scripts/user/create_example_job.sh

# 查看作业状态。
kubectl get kj
```

:::{tip}

如果希望体验隐语白屏功能([隐语白屏使用手册官方文档](https://www.secretflow.org.cn/docs/quickstart/mvp-platform))，请使用如下命令完成部署。

```bash
# 启动集群，会拉起 4 个 docker 容器，包括一个平台页面容器、一个控制平面 master 、两个 Lite 节点 alice 和 bob。
./start_standalone.sh center -u web

```

:::

{#p2p-network-mode}

### 点对点组网模式

```bash
# 启动集群，会拉起两个 docker 容器，分别表示 Autonomy 节点 alice 和 bob。
./start_standalone.sh p2p

# 登入 alice 节点容器（或 bob 节点容器）。
docker exec -it ${USER}-kuscia-autonomy-alice bash

# 创建并启动作业（两方 PSI 任务）。
scripts/user/create_example_job.sh

# 查看作业状态。
kubectl get kj
```

## 作业状态

如果作业执行成功，则 `kubectl get kj` 命令会显示类似下方的输出，Succeeded 表示成功状态：

```bash
NAME                             STARTTIME   COMPLETIONTIME   LASTRECONCILETIME   PHASE
secretflow-task-20230406162606   50s         50s              50s                 Succeeded
```

同时，在 alice 和 bob 节点容器中能看到 PSI 结果输出文件：

```bash
# 以中心化集群模式下的 alice 节点为例：
docker exec -it ${USER}-kuscia-lite-alice cat var/storage/data/psi-output.csv

# 以点对点集群模式下的 alice 节点为例：
docker exec -it ${USER}-kuscia-autonomy-alice cat var/storage/data/psi-output.csv
```

结果输出（仅前 4 行）：

```bash
id1,age,education,default,balance,housing,loan,day,duration,campaign,pdays,previous,job_blue-collar,job_entrepreneur,job_housemaid,job_management,job_retired,job_self-employed,job_services,job_student,job_technician,job_unemployed,marital_divorced,marital_married,marital_single
0,1.5306293518221676,-0.3053083468611561,-0.117184991347747,0.2145303545250443,1.0358211226635177,-0.3925867711542392,-1.2618906002715358,1.9048694929309795,-0.5762472500554522,-0.4852053503766987,-0.3619838367558999,-0.4639325546169564,-0.1731690076375218,-0.1580237499348341,1.3543943126559297,-0.2734046609851663,-0.1960131708137989,-0.3006459829345367,-0.1700475343179241,-0.4466166954522901,-0.1840186845246444,-0.3589389310523966,0.8806950470683438,-0.6902303314457872
1,1.2763683978477116,-0.3053083468611561,-0.117184991347747,-0.4786597597189064,-0.9654176557324816,-0.3925867711542392,-1.2618906002715358,3.1181827517827982,-0.5762472500554522,-0.4852053503766987,-0.3619838367558999,-0.4639325546169564,-0.1731690076375218,-0.1580237499348341,1.3543943126559297,-0.2734046609851663,-0.1960131708137989,-0.3006459829345367,-0.1700475343179241,-0.4466166954522901,-0.1840186845246444,-0.3589389310523966,0.8806950470683438,-0.6902303314457872
10,-0.2491973259990245,-0.3053083468611561,-0.117184991347747,-0.4620690476721626,1.0358211226635177,-0.3925867711542392,-1.0230801579932494,1.1740266828931782,-0.5762472500554522,-0.4852053503766987,-0.3619838367558999,-0.4639325546169564,-0.1731690076375218,-0.1580237499348341,1.3543943126559297,-0.2734046609851663,-0.1960131708137989,-0.3006459829345367,-0.1700475343179241,-0.4466166954522901,-0.1840186845246444,-0.3589389310523966,-1.1354668149080638,1.4487917358041271
```

## 停止/卸载体验集群

{#stop}

### 停止体验集群

如果你需要停止并卸载体验集群，可以直接运行[卸载脚本](#uninstall)。

获取 Kuscia 停止脚本，脚本会下载到当前目录：

```bash
docker run --rm --pull always $KUSCIA_IMAGE cat /home/kuscia/scripts/deploy/stop.sh > stop.sh && chmod u+x stop.sh
```

使用方法：

```bash
./stop.sh [center/p2p/all]

# 停止点对点组网模式集群
./stop.sh p2p

# 停止中心化组网模式集群
./stop.sh center

# 停止所有组网模式集群（可省略参数）
./stop.sh all
```

{#uninstall}

### 卸载体验集群

获取 Kuscia 卸载脚本，脚本会下载到当前目录：

```bash
docker run --rm --pull always $KUSCIA_IMAGE cat /home/kuscia/scripts/deploy/uninstall.sh > uninstall.sh && chmod u+x uninstall.sh
```

与[停止脚本](#stop)使用方法相同，运行卸载脚本将卸载相应组网模式的集群，包括删除 Kuscia 容器、volume 和 network（若无其他容器使用）等。例如：

```bash
# 卸载所有组网模式集群
./uninstall.sh
```

## 接下来

请继续阅读 [了解 KusciaJob][part-2] 章节，来了解示例作业背后的细节。

[part-2]: ./run_secretflow_cn.md
