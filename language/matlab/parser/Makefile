.PHONY: init_env
init_env:
	@if [ -d .matlab ]; then \
        echo "Directory .matlab exists"; \
    else \
        echo "Directory .matlab does not exist"; \
        echo "Init matlab python env"; \
        python3 -m venv .matlab; \
        source .matlab/bin/activate; \
        pip install -r requirements.txt; \
    fi

override VS ?= v5

.PHONY: run
run:
	@echo "parser matlab file version: ${VS}"
	@. .matlab/bin/activate && source /etc/profile && python  main.py --input ../practise/testcase/mock_33_data.mat --output ./data.csv  --field="mock_data"


help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m \033[31m%s\033[0m\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help