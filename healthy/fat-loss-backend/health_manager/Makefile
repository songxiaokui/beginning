GIT_COMMIT=$(shell git log -n 1 --pretty=format:"%H" | cut -c 1-8)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
DATE=$(shell date +"%Y%m%d%H")
VERSION?=${GIT_BRANCH}.${GIT_COMMIT}.${DATE}
DEP_TAG?=""
ARCH=arm64
WAREHOUSE_NAME := austsxk
IMAGE_NAME=fatloss_backend
IMAGE=${WAREHOUSE_NAME}/${IMAGE_NAME}_${ARCH}:${VERSION}-2

UNAME_S := $(shell uname)

ifeq ($(UNAME_S), Darwin)
    SED_INPLACE := sed -i ""
else
    SED_INPLACE := sed -i
endif

.PHONY: build
build:  ## build image
	@pip uninstall backports.zoneinfo
	@pip freeze > requirements.txt
	@echo "going to build image: ${IMAGE}"
	@docker build . -t ${IMAGE} --platform=linux/${ARCH}
	#@#docker push ${IMAGE}


.PHONY: start collectstatic run run_container format
start:
	@echo "Starting server..."
	@pip install backports.zoneinfo
	@python manage.py runserver 0.0.0.0:5011

collectstatic: ## 收集静态文件
	@python manage.py collectstatic


run:  ## 生产环境运行
	@gunicorn -c gunicorn.conf.py health_manager.wsgi:application


run_container:  ## 容器化运行
	@echo "容器化运行"
	@docker-compose up

format:  ## 格式化代码
	@echo "format code"
	@black .

.DEFAULT_GOAL: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help