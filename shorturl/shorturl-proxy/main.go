package main

import (
	"flag"
	"fmt"
	"shorturl-proxy/pkg/config"
	"shorturl-proxy/pkg/log"
	"shorturl-proxy/proxy"

	"github.com/gin-gonic/gin"
)

var (
	configFile = flag.String("config", "dev.config.yaml", "")
)

func main() {
	flag.Parse()
	config.InitConfig(*configFile)
	cnf := config.GetConfig()

	log.SetLevel(cnf.Log.Level)
	log.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	log.SetPrintCaller(true)

	logger := log.NewLogger()
	// logger.SetOutput(os.Stderr)
	logger.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	logger.SetLevel(cnf.Log.Level)
	logger.SetPrintCaller(true)

	p := proxy.NewProxy(cnf, logger)

	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {})

	public := r.Group("/p")
	public.GET("/:short_key", p.PublicProxy)
	user := r.Group("/u")
	user.GET("/:short_key", p.PublicProxy)

	r.Run(fmt.Sprintf("%s:%d", cnf.Http.IP, cnf.Http.Port))
}
