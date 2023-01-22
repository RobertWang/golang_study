
# cd ~/smallint
# go mod init smallint

go version
# go version go1.19.2 darwin/amd64


go test -bench . -benchmem ./...

# goos: darwin
# goarch: amd64
# pkg: smallint
# cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
# BenchmarkConvert-4       4970281               218.3 ns/op             0 B/op          0 allocs/op
# PASS
# ok      smallint        1.337s

