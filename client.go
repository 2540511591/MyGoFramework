package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"zeh/MyGoFramework/base"
	"zeh/MyGoFramework/conf"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:9200")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("请选择操作：\n1、登录\n2、用户信息\n")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var data []byte
		if input == "1" {
			data = login()
			fmt.Println("正在登录")
		} else if input == "2" {
			data = userInfo()
			fmt.Println("正在获取用户信息")
		} else {
			fmt.Println("无效的选项！")
			continue
		}

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println(err)
		}

		data = make([]byte, conf.ServerConfig.DataSize)
		_, err = conn.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		id, buffer := base.TestUnPack(data)

		fmt.Printf("服务器响应，id:%d,msg:%s\n", id, string(buffer))
	}
}

func login() []byte {
	data, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "123456",
	})
	return base.TestPack(1, data)
}

func userInfo() []byte {
	return base.TestPack(2, []byte{})
}
