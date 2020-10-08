package encode

import "encoding/base64"

// waitingForEncodeUrl 待编码URL
type waitingForEncodeUrl struct {
	url string
}

func FromUrl(s string) *waitingForEncodeUrl {
	return &waitingForEncodeUrl{url: s}
}

// ByBase64 对URL进行base64编码
func (s *waitingForEncodeUrl) ByBase64() string {
	return base64.URLEncoding.EncodeToString([]byte(s.url))
}

// ByBase64 对字符串进行base64编码
func (s *waitingForEncodeString) ByBase64() string {
	return base64.StdEncoding.EncodeToString([]byte(s.value))
}
