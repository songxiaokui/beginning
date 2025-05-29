.PHONY: run
run:  ## 本地运行
	@echo "access http://localhost:5173"
	@npm run dev

.DEFAULT_GOAL: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help