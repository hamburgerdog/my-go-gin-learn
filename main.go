package main

import (
	"fmt"
	"net/http"

	"xjosiah.com/go-gin/pkg/setting"
	"xjosiah.com/go-gin/routers"
)

func main() {
	router := routers.InitRouter()

	//	& 为了改变参数的值，ListenAndServer 需要用到&http.Server中的参数
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	
	s.ListenAndServe()
}
