package encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptString_ByRC4(t *testing.T) {
	cipher := NewCipher()
	cipher.SetSecretKey("123456")
	r := FromString("12345678").ByRC4(cipher).ToBase64String()
	fmt.Println("string rc4:", r)
}
