package utils

import (
	"Frame/frame/iface"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type GlobalObj struct {
	TcpServer      iface.IServer // 连接Socket对象
	Name           string        // 主机服务名
	Host           string        // 监听IP
	IPVersion      string        // 连接方式
	TcpPort        uint16        // 监听端口
	Version        string        // 框架版本号
	MaxConn        int           // 最大连接数
	MaxPackageSize uint16        // 数据包字节最大值
}

// GlobalObject 对外全局对象
var GlobalObject *GlobalObj

// LoadConf 读取、解析配置文件内的参数
func (g *GlobalObj) LoadConf() {
	data, err := ioutil.ReadFile("../conf/config.json")
	if err != nil {
		fmt.Println("Load Config error, ", err)
		panic(err)
	}
	if err := json.Unmarshal(data, &GlobalObject); err != nil {
		fmt.Println("Unmarshal Config file error, ", err)
		panic(err)
	}
}

// init 初始化方法，未加载的配置项使用的默认值。执行过后再进行加载及配置操作
func init() {
	GlobalObject = &GlobalObj{
		Name:           "Default Server",
		IPVersion:      "tcp4",
		Host:           "0.0.0.0",
		TcpPort:        8000,
		Version:        "v0.7",
		MaxConn:        10,
		MaxPackageSize: 512,
	}
	GlobalObject.LoadConf()
}
