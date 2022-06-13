# Week08 的作业


现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。

作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

完整的 httpserver-deployment.yaml 文件如下：

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: httpserver
  namespace: default
  labels:
    app: httpserver
spec:
  replicas: 4
  selector:
    matchLabels:
      app: httpserver
  template:
    metadata:
      labels:
        app: httpserver
    spec:
      containers:
        - name: httpserver
          image: miraclewong/httpserver:v8
          command: [/httpserver]
          ports:
          - containerPort: 80
          resources:
            limits:
              cpu: 500m
              memory: 512Mi
            requests:
              cpu: 500m
              memory: 512Mi
          livenessProbe:
            httpGet:
              path: /healthz
              port: 80
            initialDelaySeconds: 10
            periodSeconds: 5
          readinessProbe:
            httpGet:
              path: /healthz
              port: 80
            initialDelaySeconds: 30
            periodSeconds: 5
            successThreshold: 2
--- 
apiVersion: v1 
kind: Service 
metadata: 
  name: httpserver
spec: 
  ports: 
    - port: 80
      targetPort: 80
      nodePort: 30001 
  selector: 
    app: httpserver 
  type: NodePort

```

1. 优雅启动
这里 Dockerfile 中使用的是`ENTRYPOINT ["/httpserver"]`的方式，直接用go 编译后的二进制文件启动的，没有使用任何的类似 `start.sh`脚本的方式。
我们截取部分片段：

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: httpserver
  namespace: default
  labels:
    app: httpserver
spec:
  replicas: 4
  selector:
    matchLabels:
      app: httpserver
  template:
    metadata:
      labels:
        app: httpserver
    spec:
      containers:
        - name: httpserver
          image: miraclewong/httpserver:v2
          command: [/httpserver]
          ports:
          - containerPort: 80
```

其实和后面的探针存活的探测 `livenessProbe` 和 `readinessProbe` 一起，组成了优雅启动。

2. 优雅终止

这里可以采用preStop Hook的方式（孟老师的给的例子中，用的是nginx 镜像和相关的preStop Hook），而我们的作业中是Go服务。
在 李程远 老师的 《容器实战高手课》之 「[02 | 理解进程（1）：为什么我在容器中不能kill 1号进程？](https://time.geekbang.org/column/article/309423)」的例子中，可以看到Go的服务注册了
原因在于：
go官方文档来看，是因为 go 的 runtime，自动注册了 SIGTERM 信号，[https://pkg.go.dev/os/signal#section-directories](https://pkg.go.dev/os/signal#section-directories) 
所以，自己修改了代码，加入了相关SIGTERM（15）和 SIGKILL（9）两个信号的检测，这样在Pod 的生命周期退出时，及时没有preStop Hook的方式下，kubelet 给出SIGTERM信号后，Go 自己会收到信号进行退出的。 

```go

	// 监控两个信号
	// TERM信号（kill + 进程号 触发）
	// 中断信号（ctrl + c 触发）
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	s := <-osc
	fmt.Println("监听到退出信号,s=", s)

	// 退出前的清理操作
	// clean()

	fmt.Println("main程序退出")
	
```

3. 资源需求和 QoS 保证
这里使用的 QoS 类为Guaranteed 的方法：

```
      resources:
        limits:
          cpu: 500m
          memory: 512Mi
        requests:
          cpu: 500m
          memory: 512Mi
```

公司的微服务，大部分都是SpringCloud技术栈的Java服务，通过jar包的方法启动。
一般情况下，使用的是QoS模型中的Burstable这一类别，同时加入requests 和 limit的CPU 和 Memory的限制。

4. 探活
探活采用的是readness 和 liveness 

```
          livenessProbe:
            httpGet:
          path: /healthz
          port: 80
          initialDelaySeconds: 10
          periodSeconds: 5
            readinessProbe:
            httpGet:
          path: /healthz
          port: 80
          initialDelaySeconds: 30
          periodSeconds: 5
          successThreshold: 2
```
5. 日常运维需求，日志等级

这里main.go 的代码比较简单，没有采用具体的日志框架，因此无法实现。公司线上的服务，部分变量可以采用环境变量的方法传入到Pod中，代码读取环境变量值即可。
以下是之前线上某个服务的Charts的一部分，其中env 部分，就是引入了变量的值。日志等级的话，也可以这样实现。

```
    spec:
      containers:
        - name: jar
          image: "{{ .Values.image.war.repository }}:{{ .Values.image.war.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          env:
            - name: TZ
              value: Asia/Shanghai
```


6. 配置和代码分离

所谓的我配置和代码分离的方式，我想到的是ConfigMap 和 环境变量的方式（上面已经使用）。
由于程序中，没有使用到，这里给出的是 Kubernetes官方文档的 [ConfigMap](https://kubernetes.io/zh-cn/docs/concepts/configuration/configmap) 的示例。

下一步的话，自己可以改造下程序，将变量全部存入ConfigMap中

```
apiVersion: v1
kind: Pod
metadata:
  name: configmap-demo-pod
spec:
  containers:
    - name: demo
      image: alpine
      command: ["sleep", "3600"]
      env:
        # 定义环境变量
        - name: PLAYER_INITIAL_LIVES # 请注意这里和 ConfigMap 中的键名是不一样的
          valueFrom:
            configMapKeyRef:
              name: game-demo           # 这个值来自 ConfigMap
              key: player_initial_lives # 需要取值的键
        - name: UI_PROPERTIES_FILE_NAME
          valueFrom:
            configMapKeyRef:
              name: game-demo
              key: ui_properties_file_name
      volumeMounts:
      - name: config
        mountPath: "/config"
        readOnly: true
  volumes:
    # 你可以在 Pod 级别设置卷，然后将其挂载到 Pod 内的容器中
    - name: config
      configMap:
        # 提供你想要挂载的 ConfigMap 的名字
        name: game-demo
        # 来自 ConfigMap 的一组键，将被创建为文件
        items:
        - key: "game.properties"
          path: "game.properties"
        - key: "user-interface.properties"
          path: "user-interface.properties"
```

7. Service 的配置

```
apiVersion: v1 
kind: Service 
metadata: 
  name: httpserver
spec: 
  ports: 
    - port: 80
      targetPort: 80
      nodePort: 30001 
  selector: 
    app: httpserver 
  type: NodePort
```

说明：
- 这里通过NodePort的方式进行的曝露为端口30001，不是为了「作业二」，而是为了通过curl NodeIP:Port/healthz 的方式来判断Pods内部的httpserver 已经启动。
- 因为采用了多阶段编译的Dockerfile，里面采用了最为简单的镜像：scratch，结果内部连/bin/bash 和 /bin/sh 以及一些基础的Linux命令，如：curl 都没有。无法通过/healthz 的接口进行人为判断。
- 换为busybox的镜像，发现里面有了/bin/sh，没有/bin/bash 和 curl，虽然可以通过 kubectl exec的方式进入Pods内部，但是依旧无法通过/healthz 的接口进行人为判断。
- 在此基础上，加入了Service的NodePort曝露，采用 curl NodeIP:Port/healthz 的方式来判断内部httpserver 已经启动


