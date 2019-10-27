package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

// 获取本地相关ip
func GetIntranetIp() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("ip:", ipnet.IP.String())
			}

		}
	}
}

func TestGetIntranetIp(t *testing.T) {
	GetIntranetIp()
}

// 检测公网IP
func GetPulicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

func TestGetPulicIP(t *testing.T) {
	fmt.Println(GetPulicIP())
}