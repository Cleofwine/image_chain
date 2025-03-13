package shorturl

import (
	"mediahub/pkg/config"
	grpcclientpool "mediahub/pkg/grpc_client_pool"
	"mediahub/services"
	"sync"
)

var pool grpcclientpool.ClientPool
var once sync.Once

type shortUrlClient struct {
	services.DefaultClient
}

func GetShortUrlClientPool() grpcclientpool.ClientPool {
	once.Do(func() {
		cnf := config.GetConfig()
		client := &shortUrlClient{}
		pool = client.GetPool(cnf.DependOn.ShortUrl.Address)
	})
	return pool
}
