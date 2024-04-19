workdir:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: build
build: ## 构建容器镜像
	@echo docker login -u erubus -p $(PW)
	@echo "构建容器镜像"
	@docker build -t registry.austsxk.com/go/main-server:$(tag) -f Dockerfile .
	@docker push registry.austsxk.com/go/main-server:$(tag)

.PHONY: run
run:  ## 运行
	@echo "运行程序"
	@docker rm -f main-server
	@docker run -itd --name=main-server -v $(pwd)/static:/app/static -p 8881:8881  registry.austsxk.com/go/main-server:$(tag)
	@echo "access server: localhost:8881"

help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help