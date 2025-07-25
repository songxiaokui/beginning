GIT_COMMIT=$(shell git log -n 1 --pretty=format:"%H" | cut -c 1-8)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
DATE=$(shell date +"%Y%m%d%H")
VERSION?=${GIT_BRANCH}.${GIT_COMMIT}.${DATE}
DEP_TAG?=""
ARCH=arm64
WAREHOUSE_NAME := austsxk
IMAGE_NAME=fatloss_backend
IMAGE=${WAREHOUSE_NAME}/${IMAGE_NAME}_${ARCH}:${VERSION}

UNAME_S := $(shell uname)

ifeq ($(UNAME_S), Darwin)
    SED_INPLACE := sed -i ""
else
    SED_INPLACE := sed -i
endif

.PHONY: build
build:  ## 构建容器镜像
	@pip freeze > requirements.txt
	@echo "going to build image: ${IMAGE}"
	@docker build . -t ${IMAGE} --platform=linux/${ARCH}
	@docker push ${IMAGE}


.PHONY: start collectstatic run run_container format migrate
start:
	@echo "Starting server..."
	@python manage.py runserver 0.0.0.0:5011

collectstatic: ## 收集静态文件
	@python manage.py collectstatic


run:  ## 生产环境运行
	@gunicorn -c gunicorn.conf.py health_manager.wsgi:application


run_container:  ## 容器化运行
	@echo "容器化运行"
	@docker-compose up -d
	@echo "迁移数据库"
	@docker exec -it health-checkin-web sh -c "cd /app; python manage.py makemigrations; python manage.py migrate"

format:  ## 格式化代码
	@echo "format code"
	@black .

migrate:  ## 迁移数据库
	@python manage.py makemigrations
	@python manage.py migrate

.DEFAULT_GOAL: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help