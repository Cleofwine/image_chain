package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mediahub/pkg/config"
	"mediahub/pkg/log"
	"mediahub/pkg/storage"
	"mediahub/pkg/xerror"
	"mediahub/services"
	"mediahub/services/shorturl"
	"mediahub/services/shorturl/proto"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type File struct {
	storageFactory storage.StorageFactory
	log            log.ILogger
	config         *config.Config
}

func NewFileController(StorageFactory storage.StorageFactory, log log.ILogger, config *config.Config) *File {
	return &File{
		storageFactory: StorageFactory,
		log:            log,
		config:         config,
	}
}

func (f *File) Upload(c *gin.Context) {
	userId := c.GetInt64("user_id")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		f.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	file1, err := fileHeader.Open()
	if err != nil {
		f.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	defer file1.Close()
	content, err := io.ReadAll(file1)
	if err != nil {
		f.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	// fmt.Println(content)
	// 类型判断，仅支持jpg，gif，png
	if !isImage(io.NopCloser(bytes.NewReader(content))) {
		err = xerror.New("仅支持jpg、png、gif格式")
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	md5Digest := calMd5Degist(content)
	fileName := fmt.Sprintf("%x%s", md5Digest, path.Ext(fileHeader.Filename))
	filePath := "/public/" + fileName
	if userId != 0 {
		filePath = fmt.Sprintf("/%d/%s", userId, fileName)
	}
	s := f.storageFactory.CreateStorage()
	url, err := s.Upload(io.NopCloser(bytes.NewReader(content)), md5Digest, filePath)
	if err != nil {
		f.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	// 生成短链接
	shortPool := shorturl.GetShortUrlClientPool()
	conn := shortPool.Get()
	defer shortPool.Put(conn)
	client := proto.NewShortUrlClient(conn)
	fmt.Println(f.config.DependOn.ShortUrl.AccessToken)
	ctx := services.AppendBearerTokenToContext(context.Background(), f.config.DependOn.ShortUrl.AccessToken)
	res, err := client.GetShortUrl(ctx, &proto.Url{Url: url, IsPublic: userId == 0})
	if err != nil {
		f.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": res.Url})
}
