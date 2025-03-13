package file

import (
	"crypto/md5"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func isImage(r io.Reader) bool {
	_, _, err := image.Decode(r)
	if err != nil {
		return false
	}
	return true
}

func calMd5Degist(msg []byte) []byte {
	m := md5.New()
	m.Write(msg)
	bs := m.Sum(nil)
	return bs
}
