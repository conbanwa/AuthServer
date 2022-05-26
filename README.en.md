# Auth Server

Allows you to manage user permissions and roles

## Installation and Usage

```shell
// Adding requirements
go mod tidy
// running service
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
go.exe test -timeout 30s -run ^TestCreateUser$ authserver

go.exe test -timeout 30s -run ^TestDeleteUser$ authserver

go.exe test -timeout 30s -run ^TestCreateRole$ authserver

go.exe test -timeout 30s -run ^TestDeleteRole$ authserver
```

### Assigning users with roles

test example:

```shell
go.exe test -timeout 30s -run ^TestAddRoleToUser$ authserver
```

### Once user has been added, return token after login

### Get authorized to different roles by using token

test example:

```shell
go.exe test -timeout 30s -run ^TestAuth$ authserver
```
