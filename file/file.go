package fileutil

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/shesuyo/utility/hash"
)

func SaveToOSS(r *http.Request, field, domain string, bucket *oss.Bucket) []string {
	length := len(r.MultipartForm.File[field])
	photos := make([]string, 0, length)
	for i := 0; i < length; i++ {
		file, _ := r.MultipartForm.File[field][i].Open()
		defer file.Close()
		fileExt := strings.ToLower(filepath.Ext(r.MultipartForm.File[field][i].Filename))
		md5 := hashutil.MD5ByFile(file)
		file.Seek(0, 0)
		sha1 := hashutil.SHA1ByFile(file)
		file.Seek(0, 0)
		objectName := fmt.Sprintf("%s%s%s", md5, sha1, fileExt)
		_, err := bucket.GetObjectMeta(objectName)
		if err != nil {
			bucket.PutObject(objectName, file)
		}
		photos = append(photos, domain+objectName)
	}
	return photos
}
