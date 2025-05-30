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
build:  ## build image
	@echo "going to build image: ${IMAGE}"
	@docker build . -t ${IMAGE} --platform=linux/${ARCH}
	@docker push ${IMAGE}

.DEFAULT_GOAL: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help