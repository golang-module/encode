package encrypt

import (
	"fmt"
	"testing"
)

func TestMd5_FromString(t *testing.T) {
	r1 := FromString("12345678").ByMd5().ToHexString()
	fmt.Println("string md5:", r1)
	r2 := FromByte([]byte{0x01, 0x02, 0x03, 0x04}).ByMd5().ToHexString()
	fmt.Println("byte md5:", r2)
}
