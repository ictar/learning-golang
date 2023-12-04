对于完整的项目，必须创建目录并初始化：
```sh
$ mkdir hello
$ cd hello

$ go mod init hello
go: creating new go.mod: module hello

$ touch main.go && vim main.go
$ go build -x # 编译
```