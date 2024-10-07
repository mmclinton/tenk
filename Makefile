.DEFAULT_GOAL := install

ifeq ($(OS), Windows_NT)
    CONFIG_PATH := $(USERPROFILE)\.config\tenk
    MKDIR := mkdir
else
    CONFIG_PATH := $(HOME)/.config/tenk
    MKDIR := mkdir -p
endif

install: 
	@$(MKDIR) $(CONFIG_PATH)
	@echo '{"api_key": "$(api_key)"}' > $(CONFIG_PATH)/config.json
	@go install ./cmd/tenk
