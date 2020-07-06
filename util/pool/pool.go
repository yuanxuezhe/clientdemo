package pool

import (
	"gitee.com/yuanxuezhe/ynet/yconnpool"
	"net"
	"time"
)

//var Mysqlpool *yconnpool.ConnPool
var Connpool *yconnpool.ConnPool

func init() {
	//Mysqlpool, _ = yconnpool.NewConnPool(func() (yconnpool.ConnRes, error) {
	//	return sql.Open("mysql", "root:1@tcp(192.168.0.2:3306)/dante?parseTime=true")
	//}, 100, time.Second*100)
	//conn, err := net.Dial("tcp", "127.0.0.1:9001")
	Connpool, _ = yconnpool.NewConnPool(func() (yconnpool.ConnRes, error) {
		return net.Dial("tcp", "127.0.0.1:9001")
	}, 50, time.Second*100)
}
