# VERSION := $(shell echo $$(git describe --tags))
VERSION := $(shell echo "0.0.0")
COLOR_RED := $(shell echo -e "\033[0;31m")
COLOR_YELLOW := $(shell echo -e "\033[0;33m")
COLOR_END := $(shell echo -e "\033[0m")

list:
	@echo -e "$(COLOR_YELLOW)List of actions that can be executed $(COLOR_END)"
	@echo " -> clean"
	@echo " -> compile"
	@echo " -> run-webserver"
	@echo " -> install"
	@echo " -> fresh-install"
	@echo " -> uninstall"

clean:
	@echo -e "$(COLOR_YELLOW)Cleaning local files.$(COLOR_END)"
	rm -vrf bin || true
	rm -v main || true
	rm -v ncm-*.tar.gz || true
	@echo -e "$(COLOR_YELLOW)All local ( ./bin main ncm-*.tar.gz ) were cleaned.$(COLOR_END)"

compile: clean
	@echo -e "$(COLOR_YELLOW)Starting compiling$(COLOR_END)"
	@if [[ ! -d ./bin ]]; then mkdir ./bin; fi
	@echo -e "$(COLOR_YELLOW)Compiling for amd64$(COLOR_END)"
	@ GOOS=linux GOARCH=amd64 go build -o bin/ncm-linux ncm.go
	@echo -e "$(COLOR_YELLOW)Compiling for arm$(COLOR_END)"
	@ GOOS=linux GOARCH=arm go build -o bin/ncm-linux-arm ncm.go
	@echo -e "$(COLOR_YELLOW)Compiling for arm64$(COLOR_END)"
	@ GOOS=linux GOARCH=arm64 go build -o bin/ncm-linux-arm64 ncm.go
	# echo -e "$(COLOR_YELLOW)Compiling for 386$(COLOR_END)"
    # GOOS=freebsd GOARCH=386 go build -o bin/ncm-freebsd-386 ncm.go
	@echo -e "$(COLOR_YELLOW)Compiling for UI$(COLOR_END)"
	cd ui && ng build && cd ..
	mv ./ui/dist bin/www
	@echo -e "$(COLOR_YELLOW)Creating tar file named:  ncm-$(VERSION).tar.gz $(COLOR_END)"
	tar -cvzf ncm-$(VERSION).tar.gz bin nagios-conf-manager.service

run-webserver: clean compile
	chmod +x bin/ncm-linux
	export GIN_MODE=release && bin/ncm-linux web

fresh-install: uninstall install

install: compile
	@echo -e "$(COLOR_YELLOW)Creating directory for working files.$(COLOR_END)"
	sudo mkdir -p /var/lib/nagios-conf-manager
	sudo chown -R ${USER}:${USER} /var/lib/nagios-conf-manager
	@echo -e "$(COLOR_YELLOW)Exporting tar files to new home.$(COLOR_END)"
	tar -C /var/lib/nagios-conf-manager -xvzf ncm-$(VERSION).tar.gz
	@echo -e "$(COLOR_YELLOW)Changing running user to current logged user.$(COLOR_END)"
	sed -i "s/##USER##/${USER}/g" /var/lib/nagios-conf-manager/nagios-conf-manager.service nagios-conf-manager.service
	@echo -e "$(COLOR_YELLOW)Applying systemctl configs.$(COLOR_END)"
	sudo su root -c 'cd /lib/systemd/system/ && ln -nsf /var/lib/nagios-conf-manager/nagios-conf-manager.service nagios-conf-manager.service'
	sudo systemctl daemon-reload

uninstall:
	@echo -e "$(COLOR_YELLOW)Stopping nagios-conf-manager service.$(COLOR_END)"
	sudo systemctl disable nagios-conf-manager || true
	sudo systemctl stop nagios-conf-manager || true
	@echo -e "$(COLOR_YELLOW)Removing project files.$(COLOR_END)"
	sudo unlink /lib/systemd/system/nagios-conf-manager.service || true
	sudo rm -rf /var/lib/nagios-conf-manager || true
	sudo systemctl daemon-reload
	@echo -e "$(COLOR_YELLOW)DONE uninstall the project.$(COLOR_END)"

release-draft:
	# @SEE https://hub.github.com/hub-release.1.html
	@echo "Release a draft to git hub WIP"

all: clean compile
