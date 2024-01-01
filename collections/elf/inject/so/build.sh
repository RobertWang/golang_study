go mod init inject/so
go mod tidy
# go build -o ../inject.so .
go build -buildmode=plugin -o ./funcs.so ./funcs.go
