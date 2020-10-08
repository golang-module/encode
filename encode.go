package encode

type object struct{}

// waitingForEncodeStruct 待加密对象
type waitingForEncodeStruct struct {
	object
}

// waitingForEncodeString 待加密字符串
type waitingForEncodeString struct {
	value string
}

// waitingForEncodeFile 待加密文件
type waitingForEncodeFile struct {
	path string
}

// Todo From
func From(v interface{}) {

}

func FromString(s string) *waitingForEncodeString {
	return &waitingForEncodeString{value: s}
}

func FromStruct(s struct{}) *waitingForEncodeStruct {
	return &waitingForEncodeStruct{s}
}

func FromFile(f string) *waitingForEncodeFile {
	return &waitingForEncodeFile{path: f}
}
