### Golang基础服务器框架v0.6
##### 多路由
##### 新增消息管理模块，支持多路由业务api调度管理
##### 新增路由Map映射表（ID、路由方法），根据ID索引方法，向路由内添加方法，执行方法（需实现Pre、Post处理）
##### 修改Server模块、Connection模块Route属性，替换为MsgHandler属性
##### Connection所调度StartReader方法也许替换
##### App中添加多条路由，根据路由ID指定触发方法，及返回值