package encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptString_BySha256(t *testing.T) {
	r := FromString("12345678").BySha256().ToString()
	fmt.Println("string sha256:", r)
}
