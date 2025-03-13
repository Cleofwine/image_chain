package server

import (
	"context"
	"fmt"
	"shorturl-grpc/pkg/config"
	"shorturl-grpc/pkg/log"
	"shorturl-grpc/pkg/utils"
	"shorturl-grpc/pkg/xerror"
	"shorturl-grpc/proto"
	"shorturl-grpc/shorturl-server/cache"
	"shorturl-grpc/shorturl-server/data"
	"strconv"
	"time"
)

type shortUrlService struct {
	proto.UnimplementedShortUrlServer
	config            *config.Config
	log               log.ILogger
	cacheFactory      cache.CacheFactory
	urlMapDataFactory data.IUrlMapDataFactory
}

func NewService(config *config.Config,
	log log.ILogger,
	cacheFactory cache.CacheFactory,
	urlMapDataFactory data.IUrlMapDataFactory) proto.ShortUrlServer {
	return &shortUrlService{
		config:            config,
		log:               log,
		cacheFactory:      cacheFactory,
		urlMapDataFactory: urlMapDataFactory,
	}
}

func (s *shortUrlService) GetShortUrl(ctx context.Context, in *proto.Url) (out *proto.Url, err error) {
	if in.Url == "" {
		err := xerror.New("参数检查失败")
		s.log.Error(err)
		return nil, err
	}
	if !utils.IsUrl(in.Url) {
		err := xerror.New("参数检查失败")
		s.log.Error(err)
		return nil, err
	}

	// 根据长链接，查询数据库是否已经存在记录
	data := s.urlMapDataFactory.NewUrlMapData(in.IsPublic)
	entity, err := data.GetByOriginal(in.Url)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	if entity.ShortKey == "" {
		// 新增记录
		id, err := data.GenerateID(time.Now().Unix())
		if err != nil {
			s.log.Error(err)
			return nil, err
		}
		// 实现六十二进制
		entity.ShortKey = utils.ToBase62(id)
		entity.OriginalUrl = in.Url
		entity.ID = id
		entity.UpdateAt = time.Now().Unix()
		err = data.Update(entity)
		if err != nil {
			s.log.Error(err)
			return nil, err
		}
	}
	keyPrefix := ""
	domain := s.config.ShortDomain
	if !in.IsPublic {
		keyPrefix = "user_"
		domain = s.config.UserShortDomain
	}
	KVcache := s.cacheFactory.NewKVCache()
	defer KVcache.Destory()
	key := keyPrefix + entity.ShortKey
	err = KVcache.Set(key, entity.OriginalUrl, cache.DefaultTTL)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &proto.Url{
		Url:      domain + entity.ShortKey,
		IsPublic: in.IsPublic,
	}, nil
}

func (s *shortUrlService) GetOriginalUrl(ctx context.Context, in *proto.ShortKey) (out *proto.Url, err error) {
	if in.Key == "" {
		err := xerror.New("参数检查失败")
		s.log.Error(err)
		return nil, err
	}
	id := utils.ToBase10(in.Key)
	if id == 0 {
		err := xerror.New("参数检查失败")
		s.log.Error(err)
		return nil, err
	}
	keyPrefix := ""
	if !in.IsPublic {
		keyPrefix = "user_"
	}
	key := keyPrefix + in.Key
	kvcache := s.cacheFactory.NewKVCache()
	defer kvcache.Destory()
	originalUrl, err := kvcache.Get(key)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	if originalUrl == "" {
		// 添加一个过滤器，防止恶意访问不存在的数据，导致缓存穿透
		// 直接根据id是否小于等于最大的数据库id来判断
		// 基于缓存判断
		err = s.idFilter(id, kvcache, in.IsPublic)
		if err != nil {
			s.log.Error(err)
			return nil, err
		}
	}
	data := s.urlMapDataFactory.NewUrlMapData(in.IsPublic)
	if originalUrl == "" {
		entity, err := data.GetByID(id)
		if err != nil {
			s.log.Error(err)
			return nil, err
		}
		originalUrl = entity.OriginalUrl
	}
	// 重新设置缓存
	err = kvcache.Set(key, originalUrl, cache.DefaultTTL)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	// 累计访问次数
	err = data.IncrementTimes(id, 1, time.Now())
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	return &proto.Url{
		Url:      originalUrl,
		IsPublic: in.IsPublic,
	}, nil
}

func (s *shortUrlService) idFilter(id int64, kvCache cache.KVCache, isPublic bool) error {
	key := fmt.Sprintf("%s_%s", "url_map", "maxid")
	if !isPublic {
		key = fmt.Sprintf("%s_%s", "url_map_user", "maxid")
	}
	idStr, err := kvCache.Get(key)
	if err != nil {
		s.log.Error(err)
		return err
	}
	var res int64
	if idStr != "" {
		res, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			s.log.Error(err)
			return err
		}
	}
	if res < id {
		err = xerror.New("非法短链")
		s.log.Error(err)
		return err
	}
	return nil
}
