### Golang基础服务器框架v0.9
##### 连接管理&Hook方法
##### 管理连接个数，服务器连接数量取决于服务器内存，具体大小需测试。设计目的为保证服务器最大化运行效率
##### 当连接数量达到上限后为保证整体响应及时，应拒绝之后的连接
##### 连接管理模块内需对连接进行创建、销毁提供构造、析构方法（提供给用户Hook函数），提供管理模块初始化方法
##### 框架结构：
##### request：client → request → （触发PreHandle、Handle、PostHandle）server
##### Hook：创建连接之后、销毁连接之前
##### 1. Hook方法主体、2. 钩子方法注册、3. 方法调用
##### 在Server中内置Hook，app中如果向Hook注册方法，当达到触发时机即调用。当未添加方法时则不生效