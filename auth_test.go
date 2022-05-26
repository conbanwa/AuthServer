package main

import (
	"authserver/cryptos"
	"testing"
	"time"
)

func Test_Auth(t *testing.T) {
	cryptos.GenerateRSAKey(2048)
	var err error
	err = CreateUser("bob", "123456")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = CreateRole("admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = CreateRole("op")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = DeleteRole("administrator")
	if err == nil {
		t.Fail()
	}
	err = AddRoleToUser("bob", "admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = AddRoleToUser("bob", "op")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	token, err := Authenticate("bob", "123456")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	t.Log(len(token))
	roles, err := AllRoles(token)
	t.Log(roles)
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	ok := CheckRole(token, "admin")
	if !ok {
		t.Log("error: ", err)
		t.Fail()
	}
	ok = CheckRole(token, "administrator")
	if ok {
		t.Log("error: ", err)
		t.Fail()
	}
	err = Invalidate(token)
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	ok = CheckRole(token, "admin")
	if ok {
		t.Fail()
	}
	roles, err = AllRoles(token)
	if err == nil {
		t.Log("roles: ", roles)
		t.Fail()
	}
	expiredtoken := cryptos.GenerateToken("bob", 1)
	roles, err = AllRoles(expiredtoken)
	t.Log("roles: ", roles)
	if err != nil {
		t.Fail()
	}
	time.Sleep(2 * time.Second)
	roles, err = AllRoles(expiredtoken)
	if err.Error() != "token expired" {
		t.Log("err: ", err,roles)
		t.Fail()
	}
}
