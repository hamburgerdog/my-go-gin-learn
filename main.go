package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"xjosiah.com/go-gin/models"
	"xjosiah.com/go-gin/pkg/gredis"
	"xjosiah.com/go-gin/pkg/logging"
	"xjosiah.com/go-gin/pkg/setting"
	"xjosiah.com/go-gin/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	gredis.Setup()
	logging.Setup()

	router := routers.InitRouter()

	//	& 为了改变参数的值，ListenAndServer 需要用到&http.Server中的参数
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Panicln("Server exiting")
}
