package main

import (
	"zeh/MyGoFramework/base"
	"zeh/MyGoFramework/router"
)

func main() {
	server := base.NewDefaultServer()
	router.InitRouter(server.GetRouter())
	server.Start()
}
