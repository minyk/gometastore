PWD := `pwd`

default: build

build: cleanall linux darwin windows

cleanall: clean-linux clean-darwin clean-windows
	rm -rf ./build

clean-linux:
	rm -f ./build/hmstool-linux-amd64

clean-darwin:
	rm -f ./build/hmstool-darwin-amd64

clean-windows:
	rm -f ./build/hmstool-windows-amd64.exe

linux: clean-linux
	docker run --rm -e "GO111MODULE=on" -e "GOOS=linux" -e "GOARCH=amd64" -v $(PWD):/go/src/github.com/minyk/gometastore -w /go/src/github.com/minyk/gometastore/hmstool golang:1.12 go build -ldflags="-s -w ${GO_LDFLAGS}" -v -o build/hmstool-linux-amd64

darwin: clean-darwin
	docker run --rm -e "GO111MODULE=on" -e "GOOS=darwin" -e "GOARCH=amd64" -v $(PWD):/go/src/github.com/minyk/gometastore -w /go/src/github.com/minyk/gometastore/hmstool golang:1.12 go build -ldflags="-s -w ${GO_LDFLAGS}" -v -o build/hmstool-darwin-amd64

windows: clean-windows
	docker run --rm -e "GO111MODULE=on" -e "GOOS=windows" -e "GOARCH=amd64" -v $(PWD):/go/src/github.com/minyk/gometastore -w /go/src/github.com/minyk/gometastore/hmstool golang:1.12 go build -ldflags="-s -w ${GO_LDFLAGS}" -v -o build/hmstool-windows-amd64.exe