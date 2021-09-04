# learngo

go 常用命令

go build main.go
go build *.go         main.go 引用了同一个package下的其他函数
go run main.go     编译与执行
go run *.go

go install         将编译生成的二进制执行文件安装到GOPATH路径下，在我的本机上即为/Users/penn/go/bin

go doc 打印说明文档 背后是godoc工具在执行

go doc runtime NumCPU
go doc -src runtime NumCPU


go mod init 包管理工具，自定义包
