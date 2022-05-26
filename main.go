package main

import (
	"authserver/cryptos"
)

func main() {
	cryptos.GenerateRSAKey(1234)
	//可在test文件中测试用例
}
