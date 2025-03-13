package main

import (
	"flag"
	"fmt"

	"mediahub/controller/file"
	"mediahub/middleware"
	"mediahub/pkg/config"
	"mediahub/pkg/log"
	"mediahub/pkg/storage/cos"
	fileRouter "mediahub/routers/file"
	"mediahub/routers/home"

	"github.com/gin-gonic/gin"
)

var (
	configFile = flag.String("config", "dev.config.yaml", "")
)

func main() {
	flag.Parse()
	config.InitConfig(*configFile)
	cnf := config.GetConfig()
	// 检查配置获取
	// fmt.Printf("%+v\n", cnf)

	// 测试日志封装
	// logger := log.NewLogger() // 创建一个新的对象
	// logger.Error("报错了....")
	// log.Error("报错了???") // 直接使用包方法
	// return

	log.SetLevel(cnf.Log.Level)
	log.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	log.SetPrintCaller(true)

	logger := log.NewLogger()
	// logger.SetOutput(os.Stderr)
	logger.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	logger.SetLevel(cnf.Log.Level)
	logger.SetPrintCaller(true)

	r := gin.Default()
	// r.Use(customRecoveryMiddleware())
	api := r.Group("/mediahub-api")
	api.Use(middleware.Cors())

	storageFactory := cos.NewCosStorageFactory(cnf.Cos.BucketUrl, cnf.Cos.SecretID, cnf.Cos.SecretKey, cnf.Cos.CDNDomain)
	fileController := file.NewFileController(storageFactory, logger, cnf)

	fileRouter.InitFileRouter(api, fileController)
	home.InitHomeRouter(api)
	r.GET("/health", func(ctx *gin.Context) {})
	r.Run(fmt.Sprintf("%s:%d", cnf.Http.IP, cnf.Http.Port))
}
