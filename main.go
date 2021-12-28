package main

import (
	"clientdemo/msg"
	"encoding/json"
	"fmt"
	"gitee.com/yuanxuezhe/ynet"
	cccc "gitee.com/yuanxuezhe/ynet/Conn"
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

var wg sync.WaitGroup
var wgg sync.WaitGroup

func main() {
	start := time.Now() // 获取当前时间
	for i := 0; i < 100; i++ {
		wgg.Add(1)
		go Handler(i, &wg, &wgg)
	}
	wgg.Wait()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func Handler(i int, wg *sync.WaitGroup, wgg *sync.WaitGroup) {
	//conn := ynet.NewWsclient("ws://0.0.0.0:19001")
	conn := ynet.NewTcpclient("127.0.0.1:9001")
	fmt.Println(conn.LocalAddr(), "===>", conn.RemoteAddr())

	logon := Login{
		Type:    1,                  // 登录类型 0、注册 1、登录 2、登出
		Account: "",                 // 账号 userid/phone num/email
		Phone:   1,                  // 手机号码
		Email:   "446968454@qq.com", // 邮箱
		Passwd:  "1",                // 密码
	}

	go func(commConn cccc.CommConn) {
		for {
			// 发送消息
			buff, _ := conn.ReadMsg()
			fmt.Println(string(buff))
			wg.Done()
		}
	}(conn)

	for j := 0; j < 10; j++ {
		wg.Add(1)
		logon.Phone = i*100 + j
		logon.Email = fmt.Sprintf("%d@qq.com", logon.Phone)
		fmt.Printf("NUM: %3d\n", logon.Phone)
		jsons, errs := json.Marshal(logon) //转换成JSON返回的是byte[]
		if errs != nil {
			fmt.Println(errs.Error())
		}

		jsons = msg.PackageMsg("Login", string(jsons))
		// 发送消息
		conn.WriteMsg(jsons)
	}
	wgg.Done()
}
