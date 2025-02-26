.PHONY: run_deepseek
run_deepseek:  ## 本地运行DeepSeek聊天服务端
	@go run ${CURDIR}/cmd/deepseek/main.go --config=${CURDIR}/internal/etc/default.yaml

help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m \033[31m%s\033[0m\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help