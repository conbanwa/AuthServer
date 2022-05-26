package cryptos

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

var publicBlock []byte
var privateBlock []byte

//生成RSA私钥和公钥，保存到内存中
func GenerateRSAKey(bits int) {
	//Generate函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	privateBlock = x509.MarshalPKCS1PrivateKey(privateKey)
	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	publicBlock, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
}

//RSA加密
func encrypt(plainText string) string {
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicBlock)
	if err != nil {
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plainText))
	if err != nil {
		panic(err)
	}
	//返回密文
	return string(cipherText)
}

//RSA解密
func decrypt(cipherText string) string {
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(cipherText))
	//返回明文
	return string(plainText)
}

//Reader是一个全局、共享的密码用强随机数生成器
