GIT_COMMIT=$(shell git log -n 1 --pretty=format:"%H" | cut -c 1-8)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
DATE=$(shell date +"%Y%m%d%H")
VERSION?=${GIT_BRANCH}.${GIT_COMMIT}.${DATE}-1
DEP_TAG?=""
ARCH=amd64


WAREHOUSE_NAME := austsxk
IMAGE_NAME=fatloss
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
	@${SED_INPLACE} "s/\(FRONT_SERVICE_VERSION=\)\(.*\)/\1${VERSION}/" $(CURDIR)/.env
	@cd $(CURDIR)/devops && ./deploy.sh
	@echo "deploy successfully"

.DEFAULT_GOAL: help
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help