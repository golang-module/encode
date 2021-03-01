package encrypt

import (
	"crypto/sha256"
)

// BySha224 sha224加密
func (e Encrypt) BySha224() Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	bytes := sha256.Sum224(src)
	e.dst = bytes[:]
	return e
}
