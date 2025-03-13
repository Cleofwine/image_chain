package cos

import (
	"context"
	"encoding/base64"
	"io"
	"mediahub/pkg/storage"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type cosStorageFactory struct {
	bucketUrl string
	secretId  string
	secretKey string
	cdnDomain string
}

func NewCosStorageFactory(bucketUrl string, secretId string, secretKey string, cdnDomain string) storage.StorageFactory {
	return &cosStorageFactory{
		bucketUrl: bucketUrl,
		secretId:  secretId,
		secretKey: secretKey,
		cdnDomain: cdnDomain,
	}
}

func (f *cosStorageFactory) CreateStorage() storage.Storage {
	return newCos(f.bucketUrl, f.secretId, f.secretKey, f.cdnDomain)
}

type cosStorage struct {
	bucketUrl string
	secretId  string
	secretKey string
	cdnDomain string
}

func newCos(bucketUrl string, secretId string, secretKey string, cdnDomain string) storage.Storage {
	return &cosStorage{
		bucketUrl: bucketUrl,
		secretId:  secretId,
		secretKey: secretKey,
		cdnDomain: cdnDomain,
	}
}

func (s *cosStorage) Upload(r io.Reader, md5Digest []byte, dstPath string) (urls string, err error) {
	u, err := url.Parse(s.bucketUrl)
	if err != nil {
		return "", err
	}
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 如果使用临时密钥需要填入，临时密钥生成和使用指引参见 https://cloud.tencent.com/document/product/436/14048
			SecretID:  s.secretId,
			SecretKey: s.secretKey,
		},
	})
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: s.getContentType(dstPath),
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{},
	}
	if len(md5Digest) != 0 {
		opt.ObjectPutHeaderOptions.ContentMD5 = base64.StdEncoding.EncodeToString(md5Digest)
	}
	_, err = client.Object.Put(context.Background(), dstPath, r, opt)
	if err != nil {
		return "", err
	}
	urls = s.bucketUrl + dstPath
	if s.cdnDomain != "" {
		urls = s.cdnDomain + dstPath
	}
	return urls, err
}

func (s *cosStorage) getContentType(dstPath string) string {
	ext := strings.Trim(path.Ext(dstPath), ".")
	if ext == "jpg" {
		ext = "jpeg"
	}
	return "image/" + ext
}
