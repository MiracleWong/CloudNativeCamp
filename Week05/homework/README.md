# Week05 的作业 (选学：无需提交)

## 课后练习 5.1
按照课上讲解的方法在本地构建一个单节点的基于 HTTPS 的 etcd 集群
写一条数据
查看数据细节
删除数据


## 课后练习 5.2
在 Kubernetes 集群中创建一个高可用的 etcd 集群

### 两种方法
1. 课堂中老师的方法是 etcd-operator
    [coreos/etcd-operator](https://github.com/coreos/etcd-operator)

2. 或者采用 bitnami的etcd chart的方式

    [Helm Charts to deploy Etcd in Kubernetes](https://bitnami.com/stack/etcd/helm)

### 使用bitnami的etcd charts 建立高可用的集群

1. 添加bitnami仓库，下载chart

```
helm repo add bitnami https://charts.bitnami.com/bitnami

helm repo list

helm search repo etcd 

helm pull bitnami/etcd  ## 下载etcd chart到本地

tar zxvf etcd-8.1.1.tgz  ## 解压 chart，自己可以进行学习

docker pull bitnami/etcd:3.5.4-debian-10-r0 ## 提前下载镜像
```

2. 安装语句

```
## 设置ectd 的节点为3个，不进行持久化存储

helm install etcddemo --set replicaCount=3 --set persistence.enabled=false ./etcd
```


3. 排查语句

```
kubectl get pods

kubectl describe pods etcddemo-0

kubectl logs etcddemo-0
```

4. 安装完成后

```
$ kubectl get pods
NAME                     READY   STATUS    RESTARTS      AGE
etcddemo-0               1/1     Running   1 (87s ago)   105s
etcddemo-1               1/1     Running   1 (99s ago)   105s
etcddemo-2               1/1     Running   1 (99s ago)   105s

$ kubectl get statefulset
NAME       READY   AGE
etcddemo   3/3     2m22s

$ kubectl get svc 
NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
etcddemo            ClusterIP   10.97.108.223   <none>        2379/TCP,2380/TCP   2m1s
etcddemo-headless   ClusterIP   None            <none>        2379/TCP,2380/TCP   2m1s
kubernetes          ClusterIP   10.96.0.1       <none>        443/TCP             28d

```


5. etcd功能测试

```
# 按照etcd chart中给出的方法

kubectl run etcddemo-client --restart='Never' --image docker.io/bitnami/etcd:3.5.4-debian-10-r0 --env ROOT_PASSWORD=$(kubectl get secret --namespace default etcddemo -o jsonpath="{.data.etcd-root-password}" | base64 --decode) --env ETCDCTL_ENDPOINTS="etcddemo.default.svc.cluster.local:2379" --namespace default --command -- sleep infinity

kubectl exec --namespace default -it etcddemo-client -- bash

echo $ROOT_PASSWORD

etcdctl --user root:$ROOT_PASSWORD put /message Hello
etcdctl --user root:$ROOT_PASSWORD get /message

```