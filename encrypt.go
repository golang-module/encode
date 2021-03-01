package encrypt

import (
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
)

type Encrypt struct {
	src    []byte
	dst    []byte
	Cipher *Cipher
	Error  error
}

// FromString
func FromString(s string) Encrypt {
	return Encrypt{src: []byte(s)}
}

// FromFile
func FromFile(f string) Encrypt {
	bytes, err := ioutil.ReadFile(f)
	return Encrypt{src: bytes, Error: err}
}

// FromByte
func FromByte(b []byte) Encrypt {
	return Encrypt{src: b}
}

// ToByte 输出字节
func (e Encrypt) ToByte() []byte {
	dst := e.dst
	if dst == nil {
		dst = e.src
	}
	return dst
}

// ToString 输出字符串
func (e Encrypt) ToString() string {
	dst := e.dst
	if dst == nil {
		dst = e.src
	}
	return string(dst)
}

// ToBase64String 输出base64编码字符串
func (e Encrypt) ToBase64String() string {
	dst := e.dst
	if dst == nil {
		dst = e.src
	}
	return base64.StdEncoding.EncodeToString(dst)
}

// ToHexString 输出十六进制字符串
func (e Encrypt) ToHexString() string {
	dst := e.dst
	if dst == nil {
		dst = e.src
	}
	return hex.EncodeToString(dst)
}
