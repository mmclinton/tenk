.DEFAULT_GOAL := install

ifeq ($(OS), Windows_NT)
    CONFIG_PATH := $(subst /,\\,$(USERPROFILE)\AppData\Local\tenk)
    MKDIR := mkdir
    WRITE_JSON := echo {"api_key": "$(api_key)"} > $(CONFIG_PATH)/config.json
else 
    CONFIG_PATH := $(HOME)/.config/tenk
    MKDIR := mkdir -p
    WRITE_JSON := echo '{"api_key": "$(api_key)"}' > $(CONFIG_PATH)/config.json
endif

install: 
	@$(MKDIR) $(CONFIG_PATH)
	@$(WRITE_JSON)
	@go install ./cmd/tenk