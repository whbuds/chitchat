package main

import (
	. "chitchat/config"
	. "chitchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("8080")
}

// 通过指定端口启动 Web 服务器
func startWebServer(port string) {
	// 在入口位置初始化全局配置
	config := LoadConfig()
	r := NewRouter() // 通过 router.go 中定义的路由器来分发请求

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + config.App.Address)
	err := http.ListenAndServe(config.App.Address, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
