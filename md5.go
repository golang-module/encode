package encrypt

import (
	"crypto/md5"
)

// ByMd5 通过md5加密
func (e Encrypt) ByMd5() Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	hash := md5.New()
	hash.Write(src)
	e.dst = hash.Sum(nil)
	return e
}
