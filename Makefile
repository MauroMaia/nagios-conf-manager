# VERSION := $(shell echo $$(git describe --tags))
VERSION := $(shell echo "0.0.0")

list:
	@echo "List of actions that can be executed"
	@echo " -> clean"
	@echo " -> compile"
	@echo " -> run-webserver"
	@echo " -> install"
	@echo " -> uninstall"

clean:
	rm -vrf bin || true
	rm -v main || true
	rm -v ncm-*.tar.gz || true

compile: clean
	@echo "Compiling for every OS and Platform"
	@if [[ ! -d ./bin ]]; then mkdir ./bin; fi
	@ GOOS=linux GOARCH=amd64 go build -o bin/ncf-linux main.go
	@ GOOS=linux GOARCH=arm go build -o bin/ncf-linux-arm main.go
	@ GOOS=linux GOARCH=arm64 go build -o bin/ncf-linux-arm64 main.go
    #GOOS=freebsd GOARCH=386 go build -o bin/ncf-freebsd-386 main.go
	cd ui && ng build && cd ..
	mv ./ui/dist bin/www
	@echo "Creating tar file named: " ncm-$(VERSION).tar.gz
	tar -cvzf ncm-$(VERSION).tar.gz bin nagios-conf-manager.service

run-webserver: clean compile
	chmod +x bin/ncf-linux
	export GIN_MODE=release && bin/ncf-linux web

install:
	sudo mkdir -p /var/lib/nagios-conf-manager
	sudo chown -R ${USER}:${USER} /var/lib/nagios-conf-manager
	tar -C /var/lib/nagios-conf-manager -xvzf ncm-$(VERSION).tar.gz
	sed -i "s/##USER##/${USER}/g" /var/lib/nagios-conf-manager/nagios-conf-manager.service nagios-conf-manager.service
	sudo su root -c 'cd /lib/systemd/system/ && ln -nsf /var/lib/nagios-conf-manager/nagios-conf-manager.service nagios-conf-manager.service'
	sudo systemctl daemon-reload

uninstall:
	@echo "Stopping nagios-conf-manager service"
	sudo systemctl disable nagios-conf-manager || true
	sudo systemctl stop nagios-conf-manager || true
	@echo "Removing project files"
	sudo unlink /lib/systemd/system/nagios-conf-manager.service || true
	sudo rm -rf /var/lib/nagios-conf-manager || true
	sudo systemctl daemon-reload
	@echo "DONE uninstall the project"

release-draft:
	#@SEE https://hub.github.com/hub-release.1.html
	@echo "Release a draft to git hub WIP"

all: clean compile
