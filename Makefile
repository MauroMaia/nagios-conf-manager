
list:
	@echo "List of actions that can be executed"
	@echo " -> clean"
	@echo " -> compile"
	@echo " -> run-webserver"

clean:
	rm -vrf bin || true
	rm -v main || true

compile:
	@echo "Compiling for every OS and Platform"
	@if [[ ! -d ./bin ]]; then mkdir ./bin; fi
	GOOS=linux GOARCH=amd64 go build -o bin/ncf-linux main.go
	GOOS=linux GOARCH=arm go build -o bin/ncf-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/ncf-linux-arm64 main.go
    #GOOS=freebsd GOARCH=386 go build -o bin/ncf-freebsd-386 main.go
	@ cd ui && ng build

run-webserver: clean compile
	chmod +x bin/ncf-linux
	export GIN_MODE=release && bin/ncf-linux web

all: clean compile
