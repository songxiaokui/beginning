.PHONY: __source
__source:
	@echo "source development environment"
	@. .venv/bin/activate

.PHONY: init
init:  ## 初始化开发环境
	@echo "init development environment"
	@if [ -d ".venv" ]; then \
  		echo "skip init"; \
  	else \
  	  python3 -m venv .venv; \
  	  source .venv/bin/activate; \
  	  pip install -r requirements.txt; \
  	fi

.PHONY: notebook
notebook:  __source ## 运行Jupter Notebook
	@echo "run Jupter Notebook"
	@jupyter notebook