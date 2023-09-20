package main

import (
	"bufio"
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

		fmt.Print("请输入文本: ")
		input, err := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		_, err = conn.Write(base.TestPack(0, []byte(input)))
		if err != nil {
			fmt.Println(err)
		}

		data := make([]byte, conf.ServerConfig.DataSize)
		_, err = conn.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		id, buffer := base.TestUnPack(data)

		fmt.Printf("服务器响应，id:%d,msg:%s\n", id, string(buffer))
	}
}
