# Week03 的作业


1. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化

代码详情见Week03 的 homework，dockerfile 和 dockerfile-mod 在 httpserver 路径下：

`dockerfile` 为 传统方式的`dockerfile`

```dockerfile
FROM golang:1.17.2 AS builder
ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o httpserver .

FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 80
ENTRYPOINT ["/httpserver"]
```

`dockerfile-mod` 为 `go mod`的方式的`dockerfile`

```dockerfile
FROM golang:1.17.2 AS build

WORKDIR /
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM scratch
COPY --from=build /httpserver /httpserver
EXPOSE 80
WORKDIR /
ENTRYPOINT ["./httpserver"]
```

执行代码：
```shell
docker build -t httpserver:v2 .
```
或者：

```shell
docker build -t httpserver:v2 -f dockerfile-mod .
```

2. 将镜像推送至 docker 官方镜像仓库

DockerHub 官方的账号为：miraclewong
主页地址：[miraclewong's Profile](https://hub.docker.com/u/miraclewong/starred)

```shell
docker login ## (输入账号和密码，登录DockerHub)
docker tag httpserver:v2 miraclewong/httpserver:v2
docker push miraclewong/httpserver:v2
```

推送结果：
```shell
The push refers to repository [docker.io/miraclewong/httpserver]
5a4a152bafb0: Pushed
v2: digest: sha256:68796f3496e4715e11be5a7f78533f99a0418cbee838c5ae3b11249f407ca36d size: 527
```

        
3. 通过 docker 命令本地启动 httpserver 、

```shell
docker run --name httpserver -d -p 80:80 docker.io/miraclewong/httpserver:v2
```

测试结果：
```shell
curl 127.0.0.1:80/healthz
200
```

![pprof页面](http://images.iotop.work/uPic/20220503-pprof.png)

4. 通过 nsenter 进入容器查看 IP 配置

```shell
PID=$(docker inspect --format="{{ .State.Pid }}" httpserver)
nsenter -t $PID -n ip a
```

结果为：
```shell
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
4: eth0@if5: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
