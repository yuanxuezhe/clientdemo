package msg

import (
	"clientdemo/util"
	_ "clientdemo/util"
	"encoding/json"
)

type Msg struct {
	Id   string `json:"id"`
	Mac  string `json:"mac"`
	Body string `json:"body"`
}

func PackageMsg(id string, body string) []byte {
	m := &Msg{
		Id:   id,
		Mac:  util.Mac,
		Body: body,
	}

	jsons, err := json.Marshal(m) //转换成JSON返回的是byte[]

	if err != nil {
		panic(err)
	}
	return jsons
}
