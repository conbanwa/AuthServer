package cryptos

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const SEPARATOR = " "

//将结构体转为字符串，方便加密操作
func stringify(tp TokenPlain) string {
	return strings.Join([]string{
		strconv.Itoa(int(tp.Expire)),
		strconv.Itoa(int(tp.Created)),
		tp.Username},
		SEPARATOR)
}

//将字符串还原为结构体
func unstringify(str string) (tp TokenPlain, err error) {
	arr := strings.Split(str, SEPARATOR)
	if len(arr) < 3 {
		err = errors.New("token invalid")
		return
	}
	expire, err := strconv.Atoi(arr[0])
	if err != nil {
		return
	}
	expire64 := int64(expire)
	if time.Now().Unix() > expire64 {
		err = errors.New("token expired")
		return
	}
	created, err := strconv.Atoi(arr[1])
	tp = TokenPlain{
		Username: strings.Join(arr[2:], SEPARATOR),
		Expire:   expire64,
		Created:  int64(created),
	}
	return
}
