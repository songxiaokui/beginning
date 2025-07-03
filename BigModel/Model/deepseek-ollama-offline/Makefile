.PHONY: run
run:
	@if ! docker image inspect ollama/ollama:0.9.1-rc0 > /dev/null 2>&1; then \
		echo "Loading ollama image..."; \
		docker load -i $(CURDIR)/ollama/ollama-0.9.1-rc0-arm64.tar; \
	else \
		echo "Ollama image already exists."; \
	fi
	@if ! docker image inspect ghcr.io/open-webui/open-webui:main > /dev/null 2>&1; then \
		echo "Loading Open WebUI image..."; \
		docker load -i $(CURDIR)/open-webui/open-webui-arm64.tar; \
	else \
		echo "Open WebUI image already exists."; \
	fi
	@if ! docker image inspect ghcr.io/ollama-webui/ollama-webui:main > /dev/null 2>&1; then \
    		echo "Loading Ollama WebUI image..."; \
    		docker load -i $(CURDIR)/ollama-webui/ollama-webui-arm64.tar; \
    	else \
    		echo "Open Ollama image already exists."; \
    	fi
	@docker-compose down
	@docker-compose up -d
	@echo "Access: http://localhost:3888/"
	@echo "Please settings URL: http://ollama:11434/api"
