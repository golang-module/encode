package encode

import (
	"crypto/md5"
	"fmt"
)

// ByMd5 对字符串进行md5加密
func (s *waitingForEncodeString) ByMd5() string {
	bt := md5.Sum([]byte(s.value))
	return fmt.Sprintf("%x", bt) //将[]byte转成16进制
}

// Todo ByMd5 对文件进行md5加密
func (f *waitingForEncodeFile) ByMd5() string {
	return ""
}
