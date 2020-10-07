package encode

import (
	"crypto/md5"
	"fmt"
)

type str struct {
	name string
}

type file struct {
	path string
}

func FromString(s string) str {
	return str{name: s}
}

func FromFile(f string) file {
	return file{path: f}
}

// ByMd5 对字符串进行md5加密
func (s str) ByMd5() string {
	bt := md5.Sum([]byte(s.name))
	return fmt.Sprintf("%x", bt) //将[]byte转成16进制
}

// Todo ByMd5 对文件进行md5加密
func (f file) ByMd5() string {
	return ""
}
