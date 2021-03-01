package encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptString_BySha1(t *testing.T) {
	r := FromString("12345678").BySha1().ToHexString()
	fmt.Println("string sha1:", r)
}
