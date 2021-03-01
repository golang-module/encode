# encrypt
golang加密库

## ByJson
json加密

## base64加密
```go
// 字符串加密
encrypt.FromString("12345678").ByBase64().ToByte()  //  [77 84 73 122 78 68 85 50 78 122 103 61]
encrypt.FromString("12345678").ByBase64().ToString()  // MTIzNDU2Nzg=
encrypt.FromString("12345678").ToBase64String() // MTIzNDU2Nzg=
```

## md5加密
```go
// 字符串加密
encrypt.FromString("123456").ByMd5().ToByte() 
encrypt.FromString("123456").ByMd5().ToString() 

// 文件加密
encrypt.FromFile("./demo.pdf").ByMd5().ToByte() 
encrypt.FromFile("./demo.pdf").ByMd5().ToString() 
```
## ByAes
AES加密

## ByDes
DES加密

## By3Des
3DES加密

## ByRsa
RSA加密

## BySha1
SHA1加密

## BySha224
SHA224加密

## BySha256
SHA256加密

## BySha384
SHA384加密

## BySha512
SHA512加密
