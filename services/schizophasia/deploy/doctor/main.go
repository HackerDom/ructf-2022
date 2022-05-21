package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/usernamedt/doctor-service/pkg/workerpool"
	"log"
	"net/http"

	"github.com/usernamedt/doctor-service/pkg/logging"
	"github.com/usernamedt/doctor-service/pkg/setting"
	"github.com/usernamedt/doctor-service/routers"
)

func init() {
	print("settings init...\n")
	setting.Setup()

	print("logging init...\n")
	logging.Setup()

	print("workerpool init...\n")
	workerpool.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/usernamedt/doctor-service
// @license.name MIT
// @license.url https://github.com/usernamedt/doctor-service/blob/master/LICENSE
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	////If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}
