package base

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"sync"
	"zeh/MyGoFramework/base/iface"
	"zeh/MyGoFramework/utils"
)

type TcpConnection struct {
	//连接id，有服务器分配
	id uint64
	//连接协议
	proType uint8
	//socket
	conn *net.TCPConn
	//绑定的服务器
	server iface.IServer
	//写通道
	writeChan chan iface.IMessage

	ctx    context.Context
	cancel context.CancelFunc
	//读写锁
	lock *sync.RWMutex
	//是否已关闭
	isClose bool

	//自定义属性
	property map[string]interface{}
}

func (c *TcpConnection) Start() {
	//TODO implement me
	//读写分离
	go c.startReader()
	go c.startWriter()
}

func (c *TcpConnection) Stop() {
	//TODO implement me
	c.lock.Lock()
	defer c.lock.Unlock()

	utils.Try(func() {
		c.cancel()

		if !c.isClose {
			if c.writeChan != nil {
				close(c.writeChan)
			}

			c.conn.Close()

			c.isClose = true
		}
	}, func(i interface{}) {
		fmt.Println(i.(error))
	})
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

func (c *TcpConnection) SendBuffer(bytes []byte) {
	//TODO implement me
	msg := &DefaultMessage{}
	msg.SetId(0)
	msg.SetData(bytes)
	c.writeChan <- msg
}

func (c *TcpConnection) startReader() {
	defer c.Stop()
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("--->读操作错误,err:%s\n", err)
		}
	}()

	unPack := c.server.GetUnPack()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			//获取数据，处理数据传输边界、粘包等问题
			message, err := unPack(c)
			if err != nil {
				fmt.Printf("--->解包错误,err:%s\n", err)
				return
			}

			//实例化请求对象
			req := NewRequest(message, c)

			//分配工作协程处理
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

	pack := c.server.GetPack()
	for {
		select {
		case <-c.ctx.Done():
			return
		case message := <-c.writeChan:
			//封包
			data, err := pack(message, c)
			if err != nil {
				fmt.Printf("--->封包错误,err:%s\n", err)
				return
			}

			_, err = c.conn.Write(data)
			if err != nil {
				fmt.Printf("--->数据发送错误,err:%s\n", err)
				return
			}
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
		property:  make(map[string]interface{}),
		lock:      &sync.RWMutex{},
	}
	//生成context，读写模型依赖该context关闭协程
	c.ctx, c.cancel = context.WithCancel(server.GetCtx())

	return c
}
