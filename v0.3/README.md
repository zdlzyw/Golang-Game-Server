### Golang基础服务器框架v0.3
##### 基础route模块
##### BaseRouter Router基类，以基类实现Router，使用时可根据需求继承、重写该方法
##### 实体类继承抽象类后需要实现全部抽象方法，在此通过实体类先实现形成基类，重写该方法时因为是实体类可按需重写
##### Server中在生成socket时初始化Router，Connection中原HandleAPI由Router去处理
##### Server中实例化AddRouter方法，由用户进行添加
##### app-server中对BaseRouter方法重写实现功能
