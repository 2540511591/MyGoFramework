package api

import (
	"encoding/json"
	"fmt"
	"zeh/MyGoFramework/base/iface"
)

type Login struct {
	data map[string]interface{}
}

func (l *Login) PreHandle(request iface.IRequest) {
	//TODO implement me
	l.data = make(map[string]interface{})
	err := json.Unmarshal(request.GetData(), &l.data)
	if err != nil {
		fmt.Printf("json解析错误,err:%s\n", err)
	}
}

func (l *Login) Handle(request iface.IRequest) {
	//TODO implement me
	if l.data["username"] == "admin" && l.data["password"] == "123456" {
		info := request.GetConnection().GetAllProperty()
		info["isLogin"] = true
		info["username"] = "admin"
		info["password"] = "123456"
		request.GetResponse().Send(0, []byte("登录成功"))
		return
	}
	request.GetResponse().Send(0, []byte("登录失败"))
}

func (l *Login) PostHandle(request iface.IRequest) {
	//TODO implement me
	//panic("implement me")
}
