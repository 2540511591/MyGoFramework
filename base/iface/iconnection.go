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
	GetId() uint64
	GetProtocol() uint8
	GetConnection() net.Conn
	GetWsConnection() *websocket.Conn
	GetServer() IServer
	GetWriteChan() chan<- IMessage
	GetCtx() context.Context
	GetCancel() context.CancelFunc
}
