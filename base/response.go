package base

import (
	"errors"
	"zeh/MyGoFramework/base/iface"
)

type Response struct {
	message iface.IMessage
	conn    iface.IConnection
	hasSend bool
	req     iface.IRequest
}

func (r *Response) GetMessage() iface.IMessage {
	//TODO implement me
	return r.message
}

func (r *Response) GetMessageId() uint32 {
	//TODO implement me
	return r.message.GetId()
}

func (r *Response) GetData() []byte {
	//TODO implement me
	return r.message.GetData()
}

func (r *Response) GetConnection() iface.IConnection {
	//TODO implement me
	return r.conn
}

func (r *Response) GetRequest() iface.IRequest {
	//TODO implement me
	return r.req
}

func (r *Response) Send(u uint32, bytes []byte) {
	//TODO implement me
	r.message.SetId(u)
	r.message.SetData(bytes)
	r.message.SetLen(uint32(len(bytes) + 4))
}

func (r *Response) SendBuffer(bytes []byte) {
	//TODO implement me
	r.message.SetData(bytes)
}

func (r *Response) SendMsg(message iface.IMessage) {
	//TODO implement me
	r.message = message
}

func (r *Response) Output() error {
	//TODO implement me
	if r.message == nil || len(r.message.GetData()) <= 0 {
		return errors.New("数据为空!")
	}

	r.hasSend = true
	r.conn.GetWriteChan() <- r.message
	return nil
}

func (r *Response) Refresh() {
	//TODO implement me
	r.message = &DefaultMessage{}
}

func (r *Response) HasOutput() bool {
	//TODO implement me
	return r.hasSend
}

func NewResponse(conn iface.IConnection, req iface.IRequest) iface.IResponse {
	res := &Response{
		message: &DefaultMessage{},
		conn:    conn,
		hasSend: false,
		req:     req,
	}

	return res
}
