package encrypt

import (
	"crypto/sha512"
)

// BySha512 sha512加密
func (e Encrypt) BySha512() Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	bytes := sha512.Sum512(src)
	e.dst = bytes[:]
	return e
}
