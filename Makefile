.DEFAULT_GOAL := install

install: 
	@mkdir -p $(HOME)/.config/tenk
	@echo '{"api_key": "$(api_key)"}' > $(HOME)/.config/tenk/config.json
	@go install ./cmd/tenk