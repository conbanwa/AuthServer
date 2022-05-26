package cryptos

import (
	"testing"
	"time"
)

func Test_rsa(t *testing.T) {
	GenerateRSAKey(2048)
	message := "hello world"
	//加密
	cipherText := encrypt(message)
	t.Log("加密后为：", (cipherText))
	//解密
	plainText := decrypt(cipherText)
	if plainText != message {
		t.Log("解密后为：", (plainText))
		t.Fail()
	}
}
func Test_token(t *testing.T) {
	GenerateRSAKey(1024)
	username := "admin"
	//加密
	token := GenerateToken(username, 2*3600)
	//解密
	plain, err := DecryptToken(token)
	if err != nil {
		t.Log(err, plain)
		t.Fail()
	}
	t.Log("解密后为：", (plain))
	if plain.Username != username {
		t.Log(plain.Username)
		t.Fail()
	}
	if plain.Expire <= time.Now().Unix() {
		t.Log(plain.Expire)
		t.Fail()
	}
}
