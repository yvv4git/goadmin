# Makefile

DIR := bin
BIN_SERVER := server
BIN_CLIENT := client

all: linux arm freebsd windows

lint:
	gofmt -w server/main.go
	gofmt -w server/server.go
	gofmt -w clients/linux.go
	gofmt -w clients/windows.go

linux:
	GOOS=linux GOARCH=386 go build -o $(DIR)/$(BIN_SERVER)_linux_386.bin server/main.go server/server.go
	GOOS=linux GOARCH=386 go build -o $(DIR)/$(BIN_CLIENT)_linux_386.bin clients/linux.go

arm:
	#GOOS=linux GOARCH=arm go build -o $(DIR)/$(BIN_SERVER)_linux_arm.bin server/main.go server/server.go
	GOOS=linux GOARCH=arm go build -o $(DIR)/$(BIN_CLIENT)_linux_arm.bin clients/linux.go

freebsd:
	#GOOS=freebsd GOARCH=amd64 go build -o $(DIR)/$(BIN_SERVER)_freebsd_amd64.bin server/main.go server/server.go
	GOOS=freebsd GOARCH=amd64 go build -o $(DIR)/$(BIN_CLIENT)_freebsd_amd64.bin clients/linux.go

windows:
	#GOOS=windows GOARCH=386 go build -ldflags "-s -H windowsgui" -o $(DIR)/$(BIN_SERVER)_windows_386.exe server/main.go server/server.go
	GOOS=windows GOARCH=386 go build -ldflags "-s -H windowsgui" -o $(DIR)/$(BIN_CLIENT)_windows_386.exe client/windows.go

.PHONY: clean
clean:
	rm -r $(DIR)
