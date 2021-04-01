package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

func main() {
	var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = "192.168.3.130:8500" //consul server

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return
	}
	services, metainfo, err := client.Health().Service("serverNode", "v1000", true, &api.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		logrus.Warn("error retrieving instances from Consul: %v", err)
	}
	lastIndex = metainfo.LastIndex

	addrs := map[string]struct{}{}
	for _, service := range services {
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}
}
