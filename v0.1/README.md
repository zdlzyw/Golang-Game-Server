### Golang基础服务器框架v0.1
##### 1. 项目结构：server→frame，demo→app
##### 2. 项目使用Go Module管理，不引用GOPATH等路径
```
go env -w GO111MODULE=on
```
##### 3. 通过iface定了服务器的启动（Start）、停止（Stop）、运行服务（Serve）接口方法
##### 4. Server内实现接口方法，并按照网络连接方式定义了名称、连接方式、IP及端口的结构体。
##### 5. Server端在连接成功后打印出客户端发送的buffer内容
##### 6. 使用此框架创建main方法，并运行服务器（demo.server），client中使用标准net库连接服务器，并发送字符串给服务器→获得回调后再次打印出显示内容