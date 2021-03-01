package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

// BySha512 sha512加密
func (e Encrypt) BySha256() Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	bytes := sha256.Sum256(src)
	e.dst = []byte(hex.EncodeToString(bytes[:]))
	return e
}
