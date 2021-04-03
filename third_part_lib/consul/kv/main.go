package kv

import (
	"fmt"
	"github.com/tietang/props/v3/consul"
	"github.com/tietang/props/v3/ini"
	"github.com/tietang/props/v3/kvs"
	"time"
)

func main() {
	// 加载本地配置 获取consul地址及相关配置
	file := kvs.GetCurrentFilePath("config.ini", 2)
	conf := ini.NewIniFileCompositeConfigSource(file)

	// 设置环境：开发或正式
	conf.Set("profile", "dev")
	addr := conf.GetDefault("consul.addr", "192.168.3.130")
	contexts := conf.KeyValue("consul.contexts").Strings()

	// 获取consul配置，同时合并本地配置 consulConf 此时为最全配置
	consulConf := consul.NewCompositeConsulConfigSourceByType(contexts, addr, kvs.ContentIni)
	consulConf.Add(conf)

	fmt.Println(consulConf.KeyValue("mysql.name"))
	time.Sleep(time.Second * 10)

	fmt.Println(consulConf.KeyValue("mysql.name"))

}
