package encrypt

import (
	"crypto/rc4"
)

// ByRC4 rc4加密
func (e Encrypt) ByRC4(c *Cipher) Encrypt {
	src := e.src
	if e.dst != nil {
		src = e.dst
	}
	cipher, err := rc4.NewCipher(c.secretKey)
	if err != nil {
		e.Error = err
		return e
	}
	dst := make([]byte, len(src))
	cipher.XORKeyStream(dst, src)
	e.dst = dst
	return e
}
