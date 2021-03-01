package encrypt

// 加密模式
const (
	ECB = "ecb"
	CBC = "cbc"
	CFB = "cfb"
	OFB = "ofb"
	CTR = "ctr"
)

// 填充模式
const (
	NoPadding       = "no_padding"
	PKCS5Padding    = "pkcs5_padding"
	PKCS7Padding    = "pkcs7_padding"
	ISO10126Padding = "iso10226_padding"
)

// 定义Cipher对象
type Cipher struct {
	encryptMode string // 加密模式
	paddingMode string // 填充模式
	secretKey   []byte // 密钥
	iv          []byte // 偏移向量
	publicKey   []byte // 公钥
	privateKey  []byte // 私钥
}

// NewCipher 初始化Cipher对象
func NewCipher() *Cipher {
	return &Cipher{
		paddingMode: PKCS7Padding,
	}
}

// SetEncryptMode 设置加密方式
func (c *Cipher) SetEncryptMode(mode string) {
	c.encryptMode = mode
}

// SetPaddingMode 设置填充方式
func (c *Cipher) SetPaddingMode(mode string) {
	c.paddingMode = mode
}

// SetSecretKey 设置密钥
func (c *Cipher) SetSecretKey(key interface{}) {
	switch v := key.(type) {
	case string:
		c.secretKey = []byte(v)
	case []byte:
		c.secretKey = v
	}
}

// SetIV 设置iv
func (c *Cipher) SetIV(iv interface{}) {
	switch v := iv.(type) {
	case string:
		c.iv = []byte(v)
	case []byte:
		c.iv = v
	}
}

// SetPublicKey 设置公钥
func (c *Cipher) SetPublicKey(key interface{}) {
	switch v := key.(type) {
	case string:
		c.publicKey = []byte(v)
	case []byte:
		c.publicKey = v
	}
}

// SetPrivateKey 设置私钥
func (c *Cipher) SetPrivateKey(key interface{}) {
	switch v := key.(type) {
	case string:
		c.privateKey = []byte(v)
	case []byte:
		c.privateKey = v
	}
}
