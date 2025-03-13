package main

import (
	"flag"
	"fmt"
	"net"
	"shorturl-grpc/pkg/config"
	"shorturl-grpc/pkg/db/mysql"
	"shorturl-grpc/pkg/db/redis"
	"shorturl-grpc/pkg/log"
	"shorturl-grpc/proto"
	"shorturl-grpc/shorturl-server/cache"
	"shorturl-grpc/shorturl-server/data"
	"shorturl-grpc/shorturl-server/interceptor"
	"shorturl-grpc/shorturl-server/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var (
	configFile = flag.String("config", "dev.config.yaml", "")
)

func main() {
	flag.Parse()
	config.InitConfig(*configFile)
	cnf := config.GetConfig()
	// fmt.Printf("%+v\n", cnf)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cnf.Server.IP, cnf.Server.Port))
	if err != nil {
		panic(err)
	}

	mysql.InitMysql(cnf)
	redis.InitRedisPool(cnf)

	log.SetLevel(cnf.Log.Level)
	log.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	log.SetPrintCaller(true)

	logger := log.NewLogger()
	// logger.SetOutput(os.Stderr)
	logger.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	logger.SetLevel(cnf.Log.Level)
	logger.SetPrintCaller(true)

	urlMapDataFactory := data.NewUrlMapDataFactory(logger, mysql.GetDB(), cnf)
	cacheFactory := cache.NewRedisCacheFactory(redis.GetPool())
	shortUrlServer := server.NewService(cnf, logger, cacheFactory, urlMapDataFactory)
	healthCheck := health.NewServer()

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.UnaryAuthInterceptor, interceptor.UnaryErrorInterceptor), grpc.ChainStreamInterceptor(interceptor.StreamAuthInterceptor, interceptor.StreamErrorInterceptor))
	proto.RegisterShortUrlServer(s, shortUrlServer)
	grpc_health_v1.RegisterHealthServer(s, healthCheck)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
