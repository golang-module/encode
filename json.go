package encode

import (
	"encoding/json"
)

// ByJson 对对象进行json加密
func (s *waitingForEncodeStruct) ByJson() (string, error) {
	b, e := json.Marshal(s)
	return string(b), e
}
