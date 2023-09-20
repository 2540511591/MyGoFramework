package base

import (
	"zeh/MyGoFramework/base/iface"
)

type Request struct {
	message iface.IMessage
	conn    iface.IConnection
	router  iface.IRouter
	res     iface.IResponse
}

func (r *Request) GetMessage() iface.IMessage {
	//TODO implement me
	return r.message
}

func (r *Request) GetMessageId() uint32 {
	//TODO implement me
	return r.message.GetId()
}

func (r *Request) GetData() []byte {
	//TODO implement me
	return r.message.GetData()
}

func (r *Request) GetConnection() iface.IConnection {
	//TODO implement me
	return r.conn
}

func (r *Request) BindRouter(router iface.IRouter) {
	//TODO implement me
	r.router = router
}

func (r *Request) GetResponse() iface.IResponse {
	//TODO implement me
	return r.res
}

// 业务处理
func (r *Request) Call() {
	//TODO implement me
	r.router.PreHandle(r)
	r.router.Handle(r)
	r.router.PostHandle(r)
}

func NewRequest(message iface.IMessage, conn iface.IConnection) iface.IRequest {
	req := &Request{
		message: message,
		conn:    conn,
		res:     nil,
	}
	req.res = NewResponse(req.conn, req)

	return req
}
