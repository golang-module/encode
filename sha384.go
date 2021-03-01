package encrypt

import (
	"crypto/sha512"
)

// BySha384 sha384加密
func (e Encrypt) BySha384() Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	bytes := sha512.Sum384(src)
	e.dst = bytes[:]
	return e
}
