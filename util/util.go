package util

import (
	"fmt"
	"net"
)

var Mac = GetMac()

func GetMac() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	for _, inter := range interfaces {
		return fmt.Sprint(inter.HardwareAddr)
	}
	return ""
}
