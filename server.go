package main

import (
	"zeh/MyGoFramework/api"
	"zeh/MyGoFramework/base"
)

func main() {
	server := base.NewDefaultServer()
	g := server.GetRouter()
	g.AddRouter(0, &api.Test{})
	server.Start()
}
