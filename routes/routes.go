package routes

import (
	"chitchat/handlers"
	"net/http"
)

// WebRoute 定义一个WebRoute结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// WebRoutes 声明WebRoutes 切片存放所有web路由
type WebRoutes []WebRoute

// 定义所有Web路由
var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"newThread",
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	{
		"createThread",
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},
	{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},
	{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
}
