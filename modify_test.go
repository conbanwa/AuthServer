package main

import (
	"testing"
)

func Test_AddRoleToUser(t *testing.T) {
	var err error
	err = CreateUser("bob", "123456")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = DeleteUser("bobo")
	if err != nil {
		t.Log("error: ", err)
	}
	err = CreateRole("admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = DeleteRole("administrator")
	if err != nil {
		t.Log("error: ", err)
	}
	t.Log(AddRoleToUser("222", "w222"))
	t.Log(AddRoleToUser("bob", "w222"))
	t.Log(AddRoleToUser("222", "admin"))
	err = AddRoleToUser("bob", "admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
}

func Test_CreateUser(t *testing.T) {
	err := CreateUser("admin", "123456")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
}
func Test_DeleteUser(t *testing.T) {
	var err error
	err = CreateUser("admin", "123456")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = DeleteUser("admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
}
func Test_CreateRole(t *testing.T) {
	err := CreateRole("admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
}
func Test_DeleteRole(t *testing.T) {
	var err error
	err = CreateRole("admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
	err = DeleteRole("admin")
	if err != nil {
		t.Log("error: ", err)
		t.Fail()
	}
}
