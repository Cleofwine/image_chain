package proxy

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"shorturl-proxy/pkg/config"
	"shorturl-proxy/pkg/log"
	"shorturl-proxy/services"
	"shorturl-proxy/services/shorturl"
	"shorturl-proxy/services/shorturl/proto"

	"github.com/gin-gonic/gin"
)

type proxy struct {
	config *config.Config
	log    log.ILogger
}

func NewProxy(config *config.Config, log log.ILogger) *proxy {
	return &proxy{
		config: config,
		log:    log,
	}
}

func (p *proxy) PublicProxy(ctx *gin.Context) {
	// p.proxy(ctx, true)
	p.browserRedirection(ctx, true)
}

func (p *proxy) UserProxy(ctx *gin.Context) {
	// p.proxy(ctx, false)
	p.browserRedirection(ctx, false)
}

func (p *proxy) getOriginalUrl(shortKey string, isPublic bool) (string, error) {
	shortPool := shorturl.GetShortUrlClientPool()
	conn := shortPool.Get()
	defer shortPool.Put(conn)
	client := proto.NewShortUrlClient(conn)
	ctx := services.AppendBearerTokenToContext(context.Background(), p.config.DependOn.ShortUrl.AccessToken)
	res, err := client.GetOriginalUrl(ctx, &proto.ShortKey{Key: shortKey, IsPublic: isPublic})
	if err != nil {
		p.log.Error(err)
		return "", err
	}
	return res.Url, err
}

// 方案一：反向代理
// 1. 优点：请求过程中原始url始终被隐藏，安全性较好
// 2. 缺点：整个请求过程以及数据加载都将通过代理服务器，导致代理的压力较大
func (p *proxy) proxy(ctx *gin.Context, isPublic bool) {
	shortKey := ctx.Param("short_key")
	originalUrl, err := p.getOriginalUrl(shortKey, isPublic)
	if err != nil {
		p.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	target, _ := url.Parse(originalUrl)
	rp := httputil.NewSingleHostReverseProxy(target)
	rp.Director = func(req *http.Request) {
		req.URL.Path = target.Path
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
	}
	rp.ServeHTTP(ctx.Writer, ctx.Request)
}

// 方案二：302 重定向
// 301是永久重定向，标识请求的资源永久移动到新的URL，浏览器会将原始URL的权重转移到新的URL，并且在未来访问时直接使用新的URL
// 302是临时重定向，标识请求的资源暂时移动到一个不同的地址，是暂时的，浏览器会保存原始URL，并且在未来访问时继续使用原始URL
func (p *proxy) browserRedirection(ctx *gin.Context, isPublic bool) {
	shortKey := ctx.Param("short_key")
	originalUrl, err := p.getOriginalUrl(shortKey, isPublic)
	if err != nil {
		p.log.Error(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.Redirect(http.StatusFound, originalUrl)
}
