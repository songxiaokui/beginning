.PHONY: run build

run:
	@cd ..
	@uv run storage.py


override BUILD_ARGS ?=
override DOCKER_RESP ?= "YOU REPO URL"
override PROJECT_NAME ?= "YOU REPO NAME"
ARCH ?= arm64
TARGET_ARCH ?= $(ARCH)
REPO_NAME ?= minio_endpoint_
VS ?= v1

build:
	@echo "build ${REPO_NAME}${TARGET_ARCH} container start..."
	@docker build --platform=linux/${TARGET_ARCH} -t ${DOCKER_RESP}/${PROJECT_NAME}/${REPO_NAME}${TARGET_ARCH}:${VS} --build-arg TARGET_ARCH=${TARGET_ARCH} -f Dockerfile .
	@echo "build ${REPO_NAME}${TARGET_ARCH} container finished..."
	@docker login -u songxiaokui -p Songxiaokui123 ${DOCKER_RESP}
	@docker push ${DOCKER_RESP}/${PROJECT_NAME}/${REPO_NAME}${TARGET_ARCH}:${VS}
