package base

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync/atomic"
	"time"
	"zeh/MyGoFramework/base/iface"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       uint32
	PackFunc   func(iface.IMessage, iface.IConnection) ([]byte, error)
	UnPackFunc func(iface.IConnection) (iface.IMessage, error)

	msgHandle iface.IHandle

	ctx    context.Context
	cancel context.CancelFunc

	cid uint64
}

func (s *Server) Start() {
	//TODO implement me
	s.listenTcp()
	s.handle()
}

func (s *Server) Stop() {
	//TODO implement me
	s.cancel()
}

func (s *Server) GetCtx() context.Context {
	//TODO implement me
	return s.ctx
}

func (s *Server) GetMsgHandle() iface.IHandle {
	//TODO implement me
	return s.msgHandle
}

func (s *Server) GetPack() func(iface.IMessage, iface.IConnection) ([]byte, error) {
	//TODO implement me
	return s.PackFunc
}

func (s *Server) GetUnPack() func(iface.IConnection) (iface.IMessage, error) {
	//TODO implement me
	return s.UnPackFunc
}

func (s *Server) AddRouter(u uint32, router iface.IRouter) {
	//TODO implement me
	s.msgHandle.AddHandle(u, router)
}

func (s *Server) handle() {
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			time.Sleep(time.Second * 10)
		}
	}
}

func (s *Server) listenTcp() {
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Printf("--->初始化监听地址失败,err:%s\n", err)
		s.Stop()
		return
	}

	listen, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Printf("--->初始化监听失败,err:%s\n", err)
		s.Stop()
		return
	}

	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				conn, err := listen.AcceptTCP()
				if err != nil {
					if errors.Is(err, net.ErrClosed) {
						fmt.Printf("--->监听连接已关闭,err:%s\n", err)
						s.Stop()
						return
					}

					fmt.Printf("--->监听连接错误,err:%s\n", err)
					continue
				}

				cid := atomic.AddUint64(&s.cid, 1)
				client := NewTcpConnection(s, conn, cid)

				go client.Start()
			}
		}
	}()
}

func NewDefaultServer() iface.IServer {
	server := &Server{
		Name:       "TCPServer",
		IPVersion:  "tcp4",
		IP:         "127.0.0.1",
		Port:       9200,
		PackFunc:   DefaultPack,
		UnPackFunc: DefaultUnPack,
		msgHandle:  nil,
		ctx:        nil,
		cancel:     nil,
		cid:        0,
	}
	server.ctx, server.cancel = context.WithCancel(context.Background())
	server.msgHandle = NewHandle(server)

	return server
}
