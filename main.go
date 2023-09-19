package main

import (
	"zeh/MyGoFramework/api"
	"zeh/MyGoFramework/base"
)

func main() {
	server := base.NewDefaultServer()
	server.AddRouter(0, &api.Index{})
	server.Start()
}
