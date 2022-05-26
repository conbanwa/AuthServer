package main

import (
	"authserver/cryptos"
	"errors"
)

func CreateUser(username, password string) (err error) {
	if _, ok := userList[username]; ok {
		return errors.New("user existed")
	}
	hashed, err := hashPassword(password)
	userList[username] = User{
		Username:   username,
		Password:   hashed,
		Permission: map[string]bool{}}
	return
}
func DeleteUser(username string) (err error) {
	if _, ok := userList[username]; !ok {
		return errors.New("user not found")
	}
	delete(userList, username)
	return
}

func CreateRole(rolename string) (err error) {
	if _, ok := roleList[rolename]; ok {
		return errors.New("role existed")
	}
	roleList[rolename] = Role{
		Name: rolename}
	return
}
func DeleteRole(rolename string) (err error) {
	if _, ok := roleList[rolename]; !ok {
		return errors.New("role not found")
	}
	delete(roleList, rolename)

	//also delete roles which have signed to users.
	for _, user := range userList {
		for v := range user.Permission {
			if v == rolename {
				delete(user.Permission, rolename)
			}
		}
	}
	return
}
func AddRoleToUser(username, rolename string) (err error) {
	if _, ok := userList[username]; !ok {
		return errors.New("user not found")
	}
	if _, ok := roleList[rolename]; !ok {
		return errors.New("role not found")
	}
	userList[username].Permission[rolename] = true
	return
}
func Authenticate(username, password string) (token string, err error) {
	user, ok := userList[username]
	if !ok {
		return token, errors.New("user not found")
	}
	if ok := checkPasswordHash(password, user.Password); !ok {
		return token, errors.New("password not match")
	}
	token = cryptos.GenerateToken(username, EXPIRE_HOUR*3600)
	return
}
func Invalidate(token string) (err error) {
	cryptos.DeleteToken(token)
	return
}
func CheckRole(token string, role string) bool {
	roles, err := AllRoles(token)
	if err != nil {
		return false
	}
	for _, v := range roles {
		if v.Name == role {
			return true
		}
	}
	return false
}
func AllRoles(token string) (roles []Role, err error) {
	tp, err := cryptos.DecryptToken(token)
	if err != nil {
		return
	}
	rolemap := userList[tp.Username].Permission
	for v := range rolemap {
		if r, ok := roleList[v]; ok {
			roles = append(roles, r)
		}
	}
	return
}
