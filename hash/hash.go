package hashutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

type HashFile interface {
	io.Reader
}

func MD5ByFile(file HashFile) string {
	md5h := md5.New()
	io.Copy(md5h, file)
	return hex.EncodeToString(md5h.Sum(nil))
}

func SHA1ByFile(file HashFile) string {
	sha1h := sha1.New()
	io.Copy(sha1h, file)
	return fmt.Sprintf("%x", sha1h.Sum(nil))
}

func SHA256ByFile(file HashFile) string {
	sha256h := sha256.New()
	io.Copy(sha256h, file)
	return fmt.Sprintf("%x", sha256h.Sum(nil))
}

func SHA512ByFile(file HashFile) string {
	sha512h := sha512.New()
	io.Copy(sha512h, file)
	return fmt.Sprintf("%x", sha512h.Sum(nil))
}
