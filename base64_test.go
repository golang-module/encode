package encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptString_ByBase64(t *testing.T) {
	r := FromString("12345678").ByBase64().ToString()
	fmt.Println("string base64:", r)
}
