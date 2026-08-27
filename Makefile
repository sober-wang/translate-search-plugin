APP_NAME := translate-search-plugin

GOOS ?= linux
GOARCH ?= amd64
CGO_ENABLE := 0

FIRST_DIR := $(shell find /usr/lib -type d -name dde-grand-search-daemon )
DIR_HEAD := $(shell dirname $(FIRST_DIR) | xargs basename)
BUILD_DIR := $(APP_NAME)/usr/lib/$(DIR_HEAD)/dde-grand-search-daemon/plugins/searcher
DEB_DIR := $(APP_NAME)/DEBIAN
BINARY := $(BUILD_DIR)/$(APP_NAME)

.PHONY: install
install:
	cd cmd
	bash install.sh
	

.PHONY: build
build:
	@echo $(DIR_HEAD)
	mkdir -p $(BUILD_DIR)
	CGO_ENABLE=$(CGO_ENABLE) GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(BINARY) cmd/main.go
	cp config/translate-search-plugin.conf $(BUILD_DIR)

.PHONY: clean
clean:
	rm -rf $(APP_NAME)
	rm -rf $(APP_NAME).deb


.PHONY: deb
deb: build
	mkdir -p $(DEB_DIR)
	cp config/deb-pkg/control $(DEB_DIR)
	dpkg-deb --build translate-search-plugin translate-search-plugin.deb 
	
	
