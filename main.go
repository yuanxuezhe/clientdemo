package main

import (
	"clientdemo/msg"
	"encoding/json"
	"fmt"
	"gitee.com/yuanxuezhe/ynet"
	"strconv"
	"sync"
	"time"
)

type Login struct {
	Type    int    `json:"type"`    // 登录类型 0、注册 1、登录 2、登出
	Userid  int    `json:"userid"`  // 用户名
	Account string `json:"account"` // 账号 userid/phone num/email
	Phone   int    `json:"phone"`   // 手机号码
	Email   string `json:"email"`   // 邮箱
	Passwd  string `json:"passwd"`  // 密码
}

type Goods struct {
	Goodsid   int32  `json:"goodsid"`   //编号
	Goodsname string `json:"goodsname"` //名称
	Type      int    `json:"type"`      //商品类型
	Source    string `json:"source"`    //来源
	Url       string `json:"url"`       //链接
	Imgurl    string `json:"imgurl"`    //图片链接
	Brand     int    `json:"brand"`     //品牌
	Status    int    `json:"status"`    //状态
	Date      int    `json:"date"`      //日期
	Time      int    `json:"time"`      //时间
}

var wg *sync.WaitGroup

func main() {
	wg = &sync.WaitGroup{}
	//for i := 0; ; i++ {
	wg.Add(1)
	go Handler(1, wg)
	time.Sleep(1 * time.Second)
	//}

	wg.Wait()
}

func Handler(i int, wg *sync.WaitGroup) {
	//conn := ynet.NewWsclient("ws://0.0.0.0:19001")
	conn := ynet.NewTcpclient("192.168.2.3:9201")

	//goods := Goods{
	//	Status:    0,                   // 登录类型 0、注册 1、登录 2、登出
	//}
	//
	//jsons, errs := json.Marshal(goods) //转换成JSON返回的是byte[]
	//if errs != nil {
	//	fmt.Println(errs.Error())
	//}
	//
	//jsons = msg.PackageMsg("Goods", string(jsons))

	logon := Login{
		Type:    1,                   // 登录类型 0、注册 1、登录 2、登出
		Account: "",                  // 账号 userid/phone num/email
		Phone:   18664324257,         // 手机号码
		Email:   "4469684514@qq.com", // 邮箱
		Passwd:  "ys6303618",         // 密码
	}

	jsons, errs := json.Marshal(logon) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	jsons = msg.PackageMsg("Login", string(jsons))

	// 发送消息
	conn.WriteMsg(jsons)

	buff, _ := conn.ReadMsg()
	fmt.Println("TCP:" + strconv.Itoa(i) + "  " + string(buff))

	// 发送消息
	//conn1.WriteMsg(jsons)
	//
	//buff1, _ := conn.ReadMsg()
	//fmt.Println("W S:" +strconv.Itoa(i) + "  " + string(buff1))

	wg.Done()
}
