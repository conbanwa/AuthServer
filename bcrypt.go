package main

import "golang.org/x/crypto/bcrypt"

const (
	//password parameters
	SALT         = "passwordSaltString"
	PASSWORDCOST = 14
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+SALT), PASSWORDCOST)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+SALT))
	return err == nil
}

//bcrypt是由Niels Provos和David Mazières设计的密码哈希函数，除了加盐来抵御rainbow table 攻击之外，bcrypt的一个非常重要的特征就是自适应性，可以保证加密的速度在一个特定的范围内，即使计算机的运算能力非常高，可以通过增加迭代次数的方式，使得加密速度变慢，从而可以抵御暴力搜索攻击。
