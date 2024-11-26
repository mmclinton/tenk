.DEFAULT_GOAL := install

ifeq ($(OS), Windows_NT)
    CONFIG_PATH := $(subst /,\\,$(USERPROFILE)\AppData\Local\tenk)
    MKDIR := mkdir
    WRITE_JSON := echo {"api_key": "$(api_key)"} > $(CONFIG_PATH)/config.json
else 
    CONFIG_PATH := $(HOME)/.config/tenk
    MKDIR := mkdir -p
    WRITE_JSON := echo '{"api_key": "$(api_key)"}' > $(CONFIG_PATH)/config.json
    INSTALL_CMD := mv
endif

GOBIN ?= $(shell go env GOBIN)

ifeq ($(GOBIN),)
    ifeq ($(OS), Windows_NT)
        GOBIN := $(USERPROFILE)/go/bin
    else
        GOBIN := $(HOME)/go/bin
    endif
endif

BIN_DIR := $(GOBIN)/tenk

install: 
	@$(MKDIR) $(CONFIG_PATH)
	@$(WRITE_JSON)
	@go install ./cmd/tenk
	@if [ "$(OS)" != "Windows_NT" ]; then \
	    sudo $(INSTALL_CMD) $(BIN_DIR) /usr/local/bin/; \
	fi
	@echo "Installation complete!"

