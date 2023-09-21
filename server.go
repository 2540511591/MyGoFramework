package main

import (
	"zeh/MyGoFramework/api"
	"zeh/MyGoFramework/base"
)

func main() {
	server := base.NewDefaultServer()
	r := server.GetRouter()
	g := r.Group("api")
	{
		g2 := g.Group("user")
		g2.AddRouter(0, &api.Test{})
	}
	server.Start()
}
