package base

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"zeh/MyGoFramework/base/iface"
)

type TcpConnection struct {
	id        uint64
	proType   uint8
	conn      *net.TCPConn
	server    iface.IServer
	writeChan chan iface.IMessage

	ctx    context.Context
	cancel context.CancelFunc
}

func (c *TcpConnection) Start() {
	//TODO implement me
	go c.startReader()
	go c.startWriter()
}

func (c *TcpConnection) Stop() {
	//TODO implement me
	c.cancel()
}

func (c *TcpConnection) GetId() uint64 {
	//TODO implement me
	return c.id
}

func (c *TcpConnection) GetProtocol() uint8 {
	//TODO implement me
	return c.proType
}

func (c *TcpConnection) GetConnection() net.Conn {
	//TODO implement me
	return c.conn
}

func (c *TcpConnection) GetWsConnection() *websocket.Conn {
	//TODO implement me
	return nil
}

func (c *TcpConnection) GetServer() iface.IServer {
	//TODO implement me
	return c.server
}

func (c *TcpConnection) GetWriteChan() chan<- iface.IMessage {
	//TODO implement me
	return c.writeChan
}

func (c *TcpConnection) GetCtx() context.Context {
	//TODO implement me
	return c.ctx
}

func (c *TcpConnection) GetCancel() context.CancelFunc {
	//TODO implement me
	return c.cancel
}

func (c *TcpConnection) startReader() {
	defer c.Stop()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("--->读操作错误,err:%s\n", err)
		}
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			message, err := c.server.GetUnPack()(c)
			if err != nil {
				fmt.Printf("--->解包错误,err:%s\n", err)
				c.Stop()
				return
			}

			req := NewRequest(message, c)

			c.server.GetMsgHandle().Execute(req)
		}
	}
}

func (c *TcpConnection) startWriter() {
	defer c.Stop()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("--->写操作错误,err:%s\n", err)
		}
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		case message := <-c.writeChan:
			data, err := c.server.GetPack()(message, c)
			if err != nil {
				fmt.Printf("--->封包错误,err:%s\n", err)
				return
			}

			_, err = c.conn.Write(data)
			if err != nil {
				fmt.Printf("--->数据发送错误,err:%s\n", err)
				c.Stop()
				return
			}
		default:

		}
	}
}

func NewTcpConnection(server iface.IServer, conn *net.TCPConn, cid uint64) *TcpConnection {
	c := &TcpConnection{
		id:        cid,
		proType:   iface.PROTOCOL_TCP,
		conn:      conn,
		server:    server,
		writeChan: make(chan iface.IMessage, 10),
		ctx:       nil,
		cancel:    nil,
	}
	c.ctx, c.cancel = context.WithCancel(server.GetCtx())

	return c
}
