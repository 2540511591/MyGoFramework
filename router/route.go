package router

import (
	"fmt"
	api2 "zeh/MyGoFramework/api"
	"zeh/MyGoFramework/base/iface"
)

func InitRouter(group iface.IRouterGroup) {
	api := group.Group("api")
	{
		api.AddRouter(1, &api2.Login{})
		auth := api.GroupPl("auth", authPipeline)
		{
			auth.AddRouterPl(2, &api2.UserInfo{}, userInfoPipeline)
		}
	}
}

func authPipeline(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	res := next(request)

	fmt.Println("请求权限接口")

	return res
}

func userInfoPipeline(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	if request.GetConnection().GetProperty("isLogin") == nil || !request.GetConnection().GetProperty("isLogin").(bool) {
		request.GetResponse().SendBuffer([]byte("请登录!"))
		return request.GetResponse()
	}

	return next(request)
}
