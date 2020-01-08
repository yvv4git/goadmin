# Makefile

DIR := bin
BIN_SERVER := server
BIN_CLIENT := client
HOST_PORT_SERVER := 0.0.0.0:12345
HOST_PORT_CLIENT := 127.0.0.1:12345

all: linux arm freebsd windows

lint:
	gofmt -w server/main.go
	gofmt -w server/server.go
	gofmt -w clients/linux.go
	gofmt -w clients/windows.go

linux:
	GOOS=linux GOARCH=386 go build -ldflags "-X main.hostPort=$(HOST_PORT_SERVER)" -o $(DIR)/$(BIN_SERVER)_linux_386.bin server/main.go server/server.go
	GOOS=linux GOARCH=386 go build -ldflags "-X main.hostPort=$(HOST_PORT_CLIENT)" -o $(DIR)/$(BIN_CLIENT)_linux_386.bin clients/linux.go

arm:
	#GOOS=linux GOARCH=arm go build -o $(DIR)/$(BIN_SERVER)_linux_arm.bin server/main.go server/server.go
	GOOS=linux GOARCH=arm go build -ldflags "-X main.hostPort=$(HOST_PORT_CLIENT)" -o $(DIR)/$(BIN_CLIENT)_linux_arm.bin clients/linux.go

freebsd:
	#GOOS=freebsd GOARCH=amd64 go build -o $(DIR)/$(BIN_SERVER)_freebsd_amd64.bin server/main.go server/server.go
	GOOS=freebsd GOARCH=amd64 go build -ldflags "-X main.hostPort=$(HOST_PORT_CLIENT)" -o $(DIR)/$(BIN_CLIENT)_freebsd_amd64.bin clients/linux.go

windows:
	#GOOS=windows GOARCH=386 go build -ldflags "-s -H windowsgui" -o $(DIR)/$(BIN_SERVER)_windows_386.exe server/main.go server/server.go
	GOOS=windows GOARCH=386 go build -ldflags "-s -H windowsgui" -o $(DIR)/$(BIN_CLIENT)_windows_386.exe client/windows.go

certs:
	openssl req -x509 -newkey rsa:4096 -keyout $(DIR)/client.key -out $(DIR)/client.pem -days 365 -nodes
	openssl req -x509 -newkey rsa:4096 -keyout $(DIR)/server.key -out $(DIR)/server.pem -days 365 -nodes

.PHONY: clean
clean:
	rm -r $(DIR)
