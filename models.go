package main

//data storage
//improve performance by using map structer
//if there is a concurrent operation needed, should use sync.Map instead.
var (
	userList = map[string]User{}
	roleList = map[string]Role{}
)

const EXPIRE_HOUR = 2

type User struct {
	Username   string
	Password   string
	Permission map[string]bool
	//keep permission list unique and reduce memory cost by using RoleName as key of the Permission map
}

type Role struct {
	Name string
	//other parameters
}
