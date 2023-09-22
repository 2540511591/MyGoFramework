package iface

import (
	"context"
	"github.com/gorilla/websocket"
	"net"
)

const (
	PROTOCOL_TCP uint8 = 1
	PROTOCOL_UDP uint8 = 2
	PROTOCOL_WS  uint8 = 3
)

type IConnection interface {
	Start()
	Stop()
	//获取连接id
	GetId() uint64
	//获取该连接的通信协议
	GetProtocol() uint8
	//获取自定义信息
	SetProperty(string, interface{})
	GetProperty(string) interface{}
	GetAllProperty() map[string]interface{}
	//获取传输层连接
	GetConnection() net.Conn
	//获取ws协议连接
	GetWsConnection() *websocket.Conn
	//获取绑定的服务器
	GetServer() IServer
	//获取写通道
	GetWriteChan() chan<- IMessage
	//获取该连接的content
	GetCtx() context.Context
	//获取该连接的关闭方法
	GetCancel() context.CancelFunc
	//向客户端发送二进制
	SendBuffer([]byte)
}
