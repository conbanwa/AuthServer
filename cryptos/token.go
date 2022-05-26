package cryptos

import (
	"errors"
	"time"
)

var tokenStorage = map[string]string{}

type TokenPlain struct {
	Username string
	Expire   int64
	Created  int64
}

func GenerateToken(username string, Expire time.Duration) string {
	tp := TokenPlain{
		Username: username,
		Expire:   time.Now().Add(Expire * time.Second).Unix(),
		Created:  time.Now().Unix(),
	}
	cipherText := encrypt(stringify(tp))
	tokenStorage[username] = cipherText
	return cipherText
}

func DeleteToken(token string) {
	tp, err := DecryptToken(token)
	if err != nil {
		//token invalid
		return
	}
	delete(tokenStorage, tp.Username)
	//it won't panic even the key is not exit
}

//读取token中的结构体
func DecryptToken(token string) (tp TokenPlain, err error) {
	str := decrypt(token)
	tmp, err := unstringify(str)
	if err != nil {
		return
	}
	stored, ok := tokenStorage[tmp.Username]
	if !ok {
		err = errors.New("user invalidated")
		return
	}
	if stored != token {
		err = errors.New("token out of date")
		return
	}
	tp = tmp
	return
}
