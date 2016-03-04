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

//获取string的sha1(32位)
func SHA1(str string) string {
	sha1new := sha1.New()
	sha1new.Write([]byte(str))
	sha1hash := fmt.Sprintf("%x", sha1new.Sum(nil))
	return sha1hash
}

func SHA512(str string) string {
	sha512new := sha512.New()
	sha512new.Write([]byte(str))
	sha1hash := fmt.Sprintf("%x", sha512new.Sum(nil))
	return sha1hash
}

func MD5(str string) string {
	md5new := md5.New()
	md5new.Write([]byte(str))
	md5hash := fmt.Sprintf("%x", md5new.Sum(nil))
	return md5hash
}
