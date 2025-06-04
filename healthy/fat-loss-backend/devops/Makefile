.PHONY: db
db:  ## 启动数据库服务
	@echo "running database"
	@docker-compose -f database.docker-compose.yaml up -d


.DEFAULT_GOAL: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help