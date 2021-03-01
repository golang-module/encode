package encrypt

import (
	"encoding/base64"
)

// ByBase64 通过base64编码
func (e Encrypt) ByBase64() Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	e.dst = []byte(base64.StdEncoding.EncodeToString(src))
	return e
}
