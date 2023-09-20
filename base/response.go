package base

import (
	"errors"
	"sync"
	"zeh/MyGoFramework/base/iface"
)

type Response struct {
	conn iface.IConnection
	req  iface.IRequest
	//消息列表
	messages map[uint32]iface.IMessage
	//消息状态，是否已发送
	mStatus map[uint32]bool
	//消息数量
	mNum uint32
	lock *sync.RWMutex
}

func (r *Response) GetConnection() iface.IConnection {
	//TODO implement me
	return r.conn
}

func (r *Response) GetRequest() iface.IRequest {
	//TODO implement me
	return r.req
}

func (r *Response) GetMessageNumber() uint32 {
	//TODO implement me
	r.lock.RLock()
	defer r.lock.RUnlock()

	return uint32(len(r.messages))
}

func (r *Response) GetMessage(u uint32) iface.IMessage {
	//TODO implement me
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.messages[u-1]
}

func (r *Response) Send(u uint32, bytes []byte) {
	//TODO implement me
	msg := &DefaultMessage{}
	msg.SetId(u)
	msg.SetData(bytes)
	r.SendMsg(msg)
}

func (r *Response) SendBuffer(bytes []byte) {
	//TODO implement me
	r.Send(0, bytes)
}

func (r *Response) SendMsg(message iface.IMessage) {
	//TODO implement me
	r.lock.Lock()
	defer r.lock.Unlock()

	r.conn.GetWriteChan() <- message
	r.messages[r.mNum] = message
	r.mStatus[r.mNum] = true
	r.mNum++
}

func (r *Response) WaitSend(id uint32, data []byte) uint32 {
	//TODO implement me
	msg := &DefaultMessage{}
	msg.SetId(id)
	msg.SetData(data)
	return r.WaitSendMsg(msg)
}

func (r *Response) WaitBuffer(data []byte) uint32 {
	//TODO implement me
	return r.WaitSend(0, data)
}

func (r *Response) WaitSendMsg(message iface.IMessage) uint32 {
	//TODO implement me
	r.lock.Lock()
	defer r.lock.Unlock()

	r.messages[r.mNum] = message
	r.mStatus[r.mNum] = false
	id := r.mNum
	r.mNum++

	return id
}

// 将消息全部发送
func (r *Response) Output() error {
	//TODO implement me
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.messages == nil {
		return errors.New("可发送数据为空!")
	}

	chl := r.conn.GetWriteChan()
	for k, v := range r.mStatus {
		//只输出未发送的消息
		if !v {
			chl <- r.messages[k]
			r.mStatus[k] = true
		}
	}

	return nil
}

// 清空消息
func (r *Response) Refresh() {
	//TODO implement me
	r.lock.Lock()
	defer r.lock.Unlock()

	r.messages = make(map[uint32]iface.IMessage)
	r.mStatus = make(map[uint32]bool)
	r.mNum = 0
}

func NewResponse(conn iface.IConnection, req iface.IRequest) iface.IResponse {
	res := &Response{
		conn:     conn,
		req:      req,
		lock:     &sync.RWMutex{},
		messages: make(map[uint32]iface.IMessage),
		mStatus:  make(map[uint32]bool),
		mNum:     0,
	}

	return res
}
