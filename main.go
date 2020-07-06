package main

import (
	"clientdemo/msg"
	"clientdemo/util/pool"
	"encoding/json"
	"fmt"
	network "gitee.com/yuanxuezhe/ynet/tcp"
	"net"
	"strconv"
	"sync"
	"time"
)

type Login struct {
	Type    int    `json:"type"`    // 登录类型 0、注册 1、登录 2、登出
	Userid  string `json:"userid"`  // 用户名
	Account string `json:"account"` // 账号 userid/phone num/email
	Phone   int    `json:"phone"`   // 手机号码
	Email   string `json:"email"`   // 邮箱
	Passwd  string `json:"passwd"`  // 密码
}

var wg *sync.WaitGroup

func main() {
	wg = &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Handler(1, wg)
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
}

func Handler(i int, wg *sync.WaitGroup) {
	conn, err := pool.Connpool.Get()
	if err != nil {
		panic(err)
	}

	logon := Login{
		Type:    1,                   // 登录类型 0、注册 1、登录 2、登出
		Account: "",                  // 账号 userid/phone num/email
		Phone:   186643242567,        // 手机号码
		Email:   "4469684514@qq.com", // 邮箱
		Passwd:  "ys6303618",         // 密码
	}

	jsons, errs := json.Marshal(logon) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	jsons = msg.PackageMsg("Login", string(jsons))

	// 发送消息
	network.SendMsg(conn.(net.Conn), jsons)

	buff, _ := network.ReadMsg(conn.(net.Conn))
	fmt.Println(strconv.Itoa(i) + "  " + string(buff))

	pool.Connpool.Put(conn)
	wg.Done()
}
