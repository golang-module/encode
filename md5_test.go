package encode

import (
	"fmt"
	"testing"
)

func TestMd5_FromString(t *testing.T) {
	md := FromString("123456").ByMd5()
	fmt.Println(md)
}
