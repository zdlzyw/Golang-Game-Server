### Golang基础服务器框架v0.8
##### 消息队列（MessageQueue）
##### 当前业务Reader和Writer分别为阻塞状态，不会占用CPU资源。当客户端进行连接后会通过调度切换Goroutine执行DoMsgHandler方法，造成系统资源消耗，需在此优化
##### 每个客户端都需要执行DoMsgHandler方法，和客户端数量一致，在此可以固定DoMsgHandler数量（池化）
##### 收到客户端读请求后发送给消息队列，由channel传递给worker，每个队列固定一个worker，worker组成池，处理完毕后重新返回给客户端Writer
##### 请求进入消息队列时需要由算法控制（负载均衡）,消息队列数量和worker数量由业务压力决定（需测试）,Worker池数量从全局配置文件获取
##### 每个worker用一个Go承载，并阻塞状态等待channel消息。当前Connection为每个Request调用DoMsgHandler，修改为发送Request给Worker池处理。保证每个消息队列收到的消息数量均衡
##### 对当前Request新增ID字段（自增），每次获得request后，通过ID和池ID取余，决定进入的池ID，再对应池内处理request请求，因池数量固定，固不会存在客户端连接请求过多时负担过高（每个请求都需创建Goroutine）
##### 调用方法：1. 在Listener之前开启Worker 2. Connection中Reader时直接调用Goroutine的方法使用池处理