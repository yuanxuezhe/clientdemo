package main

import (
	"clientdemo/msg"
	"encoding/json"
	"fmt"
	network "gitee.com/yuanxuezhe/ynet/tcp"
	"net"
	//"time"
)

type Login struct {
	Type    int    `json:"type"`    // 登录类型 0、注册 1、登录 2、登出
	Userid  string `json:"userid"`  // 用户名
	Account string `json:"account"` // 账号 userid/phone num/email
	Phone   int    `json:"phone"`   // 手机号码
	Email   string `json:"email"`   // 邮箱
	Passwd  string `json:"passwd"`  // 密码
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9001")
	if err != nil {
		panic(err)
	}

	logon := Login{
		Type: 0, // 登录类型 0、注册 1、登录 2、登出
		//Account:		// 账号 userid/phone num/email
		Userid: "yuan379152355",
		Phone:  18664324256,        // 手机号码
		Email:  "446968454@qq.com", // 邮箱
		Passwd: "ys6303618",        // 密码
	}

	jsons, errs := json.Marshal(logon) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	jsons = msg.PackageMsg("Login", string(jsons))

	// 发送消息
	network.SendMsg(conn, jsons)

	buff, _ := network.ReadMsg(conn)
	fmt.Println(string(buff))
}
