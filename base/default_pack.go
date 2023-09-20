package base

import (
	"encoding/binary"
	"errors"
	"zeh/MyGoFramework/base/iface"
	"zeh/MyGoFramework/conf"
)

// 前4字节为msgID，后面为数据
func DefaultPack(message iface.IMessage, conn iface.IConnection) ([]byte, error) {
	var data []byte
	switch conn.GetProtocol() {
	case iface.PROTOCOL_TCP:
		fallthrough
	case iface.PROTOCOL_UDP:
		data = make([]byte, conf.ServerConfig.DataSize)
		binary.BigEndian.PutUint32(data[:4], message.GetId())
		copy(data[4:], message.GetData())
	case iface.PROTOCOL_WS:
		//ws不固定长度
		data = make([]byte, 4)
		binary.BigEndian.PutUint32(data, message.GetId())
		data = append(data, message.GetData()...)
	default:
		return nil, errors.New("未知协议")
	}

	return data, nil
}

// 前4字节为msgID，后面为数据
func DefaultUnPack(conn iface.IConnection) (iface.IMessage, error) {
	var data []byte
	var dataSize uint32
	switch conn.GetProtocol() {
	case iface.PROTOCOL_TCP:
		fallthrough
	case iface.PROTOCOL_UDP:
		//读取固定长度
		dataSize = conf.ServerConfig.DataSize
		c := conn.GetConnection()
		data = make([]byte, dataSize)
		_, err := c.Read(data)
		if err != nil {
			conn.Stop()
			return &DefaultMessage{}, errors.New("连接已关闭")
		}
	case iface.PROTOCOL_WS:
		//ws不固定长度
		c := conn.GetWsConnection()
		var err error
		_, data, err = c.ReadMessage()
		if err != nil {
			conn.Stop()
			return &DefaultMessage{}, errors.New("连接已关闭")
		}
		dataSize = uint32(len(data))
	default:
		return &DefaultMessage{}, errors.New("未知协议")
	}

	//前4个字节为msgID，其他为数据
	id := data[:4]
	buffer := data[4:]
	message := new(DefaultMessage)
	message.SetId(binary.BigEndian.Uint32(id))
	message.SetData(buffer)
	return message, nil
}

func TestPack(id uint32, data []byte) []byte {
	newData := make([]byte, conf.ServerConfig.DataSize)
	binary.BigEndian.PutUint32(newData[:4], id)
	copy(newData[4:], data)

	return newData
}

func TestUnPack(data []byte) (uint32, []byte) {
	return binary.BigEndian.Uint32(data[:4]), data[4:]
}
