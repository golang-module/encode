package encrypt

import (
	"crypto/sha1"
)

// BySha1 sha1加密
func (e Encrypt) BySha1() Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	bytes := sha1.Sum(src)
	e.dst = bytes[:]
	return e
}
