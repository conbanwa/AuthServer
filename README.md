# AuthServer

Allows you to manage user permissions and roles

## Installation and Usage

```shell
// Adding requirments
go mod tidy
// runing service
go run main.go

//or
go build main.go
./main.exe
```

## What It Does

This package allows you to manage user permissions and roles.

### Add, remove users and roles

test example:

```shell
go.exe test -timeout 30s -run ^Test_CreateUser$ authserver

go.exe test -timeout 30s -run ^Test_DeleteUser$ authserver

go.exe test -timeout 30s -run ^Test_CreateRole$ authserver

go.exe test -timeout 30s -run ^Test_DeleteRole$ authserver
```

### Asing users with roles

test example:

```shell
go.exe test -timeout 30s -run ^Test_AddRoleToUser$ authserver
```

### Once user added, get token while login

### Using token to get authorized according to different roles

test example:

```shell
go.exe test -timeout 30s -run ^Test_Auth$ authserver
```
