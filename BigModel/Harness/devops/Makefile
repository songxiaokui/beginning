SHELL := /bin/bash

.DEFAULT_GOAL := help


# ============================================================
# Image
# ============================================================

IMAGE ?= austsxk/deepseek-harness-dev-base

TAG ?= 0.1.0-rc.6-dev

BASE_IMAGE ?= runzhliu/deepseek-harness:0.1.0-rc.6


# ============================================================
# Platforms
# ============================================================

PLATFORMS ?= linux/amd64,linux/arm64

BUILDER ?= harness-multiarch-builder


# ============================================================
# Go
# ============================================================
#
# latest:
#
# 每次 Go 安装 layer 被重新执行时，
# 自动查询官方最新 stable。
#
# GO_REFRESH 默认每天变化一次，
# 避免 Docker 永久缓存旧 Go 版本。
#
# ============================================================

GO_VERSION ?= latest

GO_REFRESH ?= $(shell date -u +%Y-%m-%d)


# ============================================================
# Build Metadata
# ============================================================

BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

VCS_REF := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)


# ============================================================
# Build Args
# ============================================================

BUILD_ARGS := \
	--build-arg BASE_IMAGE=$(BASE_IMAGE) \
	--build-arg GO_VERSION=$(GO_VERSION) \
	--build-arg GO_REFRESH=$(GO_REFRESH) \
	--build-arg BUILD_DATE=$(BUILD_DATE) \
	--build-arg VCS_REF=$(VCS_REF)


# ============================================================
# Help
# ============================================================

.PHONY: help
help:
	@echo ""
	@echo "DeepSeek Harness Development Base"
	@echo ""
	@echo "Image:"
	@echo "  $(IMAGE):$(TAG)"
	@echo ""
	@echo "Commands:"
	@echo ""
	@echo "  make builder"
	@echo "      Create / initialize Docker Buildx builder"
	@echo ""
	@echo "  make build"
	@echo "      Build image for current Docker architecture"
	@echo ""
	@echo "  make build-amd64"
	@echo "      Build linux/amd64 image locally"
	@echo ""
	@echo "  make build-arm64"
	@echo "      Build linux/arm64 image locally"
	@echo ""
	@echo "  make push"
	@echo "      Build amd64 + arm64 and push multi-arch image"
	@echo ""
	@echo "  make push-no-cache"
	@echo "      Force rebuild amd64 + arm64 and push"
	@echo ""
	@echo "  make inspect"
	@echo "      Inspect remote multi-architecture manifest"
	@echo ""
	@echo "  make test"
	@echo "      Test locally loaded image"
	@echo ""
	@echo "  make test-docker"
	@echo "      Test Docker socket access"
	@echo ""
	@echo "  make shell"
	@echo "      Enter local image shell"
	@echo ""
	@echo "  make version"
	@echo "      Show current build configuration"
	@echo ""
	@echo "  make clean-builder"
	@echo "      Remove Buildx builder"
	@echo ""


# ============================================================
# Buildx Builder
# ============================================================

.PHONY: builder
builder:
	@if ! docker buildx inspect $(BUILDER) >/dev/null 2>&1; then \
		echo "Creating Buildx builder: $(BUILDER)"; \
		docker buildx create \
			--name $(BUILDER) \
			--driver docker-container \
			--use; \
	else \
		echo "Using existing Buildx builder: $(BUILDER)"; \
		docker buildx use $(BUILDER); \
	fi
	@docker buildx inspect --bootstrap


# ============================================================
# Build Current Architecture
# ============================================================
#
# Mac M1/M2/M3/M4:
#
#   linux/arm64
#
# Normal Intel / AMD Linux:
#
#   linux/amd64
#
# ============================================================

.PHONY: build
build: builder
	@ARCH=$$(docker version --format '{{.Server.Arch}}'); \
	echo "Building platform: linux/$$ARCH"; \
	docker buildx build \
		--builder $(BUILDER) \
		--platform linux/$$ARCH \
		$(BUILD_ARGS) \
		--tag $(IMAGE):$(TAG) \
		--load \
		.


# ============================================================
# Build AMD64
# ============================================================

.PHONY: build-amd64
build-amd64: builder
	docker buildx build \
		--builder $(BUILDER) \
		--platform linux/amd64 \
		$(BUILD_ARGS) \
		--tag $(IMAGE):$(TAG)-amd64 \
		--load \
		.


# ============================================================
# Build ARM64
# ============================================================

.PHONY: build-arm64
build-arm64: builder
	docker buildx build \
		--builder $(BUILDER) \
		--platform linux/arm64 \
		$(BUILD_ARGS) \
		--tag $(IMAGE):$(TAG)-arm64 \
		--load \
		.


# ============================================================
# Multi-Architecture Build + Push
# ============================================================
#
# 最终生成：
#
# IMAGE:TAG
#   ├── linux/amd64
#   └── linux/arm64
#
# multi-platform image 无法同时 --load 到本地普通
# Docker image store，所以这里直接 --push。
#
# ============================================================

.PHONY: push
push: builder
	docker buildx build \
		--builder $(BUILDER) \
		--platform $(PLATFORMS) \
		$(BUILD_ARGS) \
		--tag $(IMAGE):$(TAG) \
		--push \
		.


# ============================================================
# Force Rebuild + Push
# ============================================================

.PHONY: push-no-cache
push-no-cache: builder
	docker buildx build \
		--builder $(BUILDER) \
		--platform $(PLATFORMS) \
		$(BUILD_ARGS) \
		--pull \
		--no-cache \
		--tag $(IMAGE):$(TAG) \
		--push \
		.


# ============================================================
# Inspect Multi Architecture Image
# ============================================================

.PHONY: inspect
inspect:
	docker buildx imagetools inspect $(IMAGE):$(TAG)


# ============================================================
# Local Environment Test
# ============================================================

.PHONY: test
test:
	docker run \
		--rm \
		--entrypoint /bin/bash \
		$(IMAGE):$(TAG) \
		-lc '\
			set -e; \
			\
			echo; \
			echo "===================="; \
			echo "Git"; \
			echo "===================="; \
			git --version; \
			\
			echo; \
			echo "===================="; \
			echo "Python"; \
			echo "===================="; \
			python --version; \
			pip3 --version; \
			\
			echo; \
			echo "===================="; \
			echo "Go"; \
			echo "===================="; \
			go version; \
			gopls version; \
			staticcheck -version; \
			dlv version; \
			\
			echo; \
			echo "===================="; \
			echo "C / C++"; \
			echo "===================="; \
			gcc --version | head -n 1; \
			g++ --version | head -n 1; \
			clang --version | head -n 1; \
			cmake --version | head -n 1; \
			ninja --version; \
			\
			echo; \
			echo "===================="; \
			echo "Docker"; \
			echo "===================="; \
			docker --version; \
			docker compose version; \
			docker buildx version; \
			\
			echo; \
			echo "Environment test passed."; \
		'


# ============================================================
# Test Docker Host Socket
# ============================================================
#
# 注意：
#
# 这一步把宿主机 Docker Socket 暴露给容器。
#
# Harness 正式运行时也是同样原理。
#
# ============================================================

.PHONY: test-docker
test-docker:
	docker run \
		--rm \
		--user 0:0 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		--entrypoint /bin/bash \
		$(IMAGE):$(TAG) \
		-lc '\
			set -e; \
			docker version; \
			echo; \
			docker ps; \
			echo; \
			docker info --format "Docker Server: {{.ServerVersion}} / {{.Architecture}}"; \
		'


# ============================================================
# Interactive Shell
# ============================================================

.PHONY: shell
shell:
	docker run \
		--rm \
		-it \
		--user 0:0 \
		--entrypoint /bin/bash \
		$(IMAGE):$(TAG)


# ============================================================
# Interactive Shell + Docker Socket
# ============================================================

.PHONY: shell-docker
shell-docker:
	docker run \
		--rm \
		-it \
		--user 0:0 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		--entrypoint /bin/bash \
		$(IMAGE):$(TAG)


# ============================================================
# Print Configuration
# ============================================================

.PHONY: version
version:
	@echo ""
	@echo "IMAGE       = $(IMAGE)"
	@echo "TAG         = $(TAG)"
	@echo "BASE_IMAGE  = $(BASE_IMAGE)"
	@echo "PLATFORMS   = $(PLATFORMS)"
	@echo "BUILDER     = $(BUILDER)"
	@echo "GO_VERSION  = $(GO_VERSION)"
	@echo "GO_REFRESH  = $(GO_REFRESH)"
	@echo "BUILD_DATE  = $(BUILD_DATE)"
	@echo "VCS_REF     = $(VCS_REF)"
	@echo ""


# ============================================================
# Clean Builder
# ============================================================

.PHONY: clean-builder
clean-builder:
	-docker buildx rm $(BUILDER)

.PHONY: start down restart


start:
	@docker compose up -d
	@docker exec deepseek-harness sh -lc 'mkdir -p "$$HOME/.dsh" && touch "$$HOME/.dsh/settings.yaml" && if grep -q "^[[:space:]]*defaultPreset:" "$$HOME/.dsh/settings.yaml"; then sed -i "s/^[[:space:]]*defaultPreset:.*/  defaultPreset: danger-full-access/" "$$HOME/.dsh/settings.yaml"; elif grep -q "^permission:" "$$HOME/.dsh/settings.yaml"; then sed -i "/^permission:/a\  defaultPreset: danger-full-access" "$$HOME/.dsh/settings.yaml"; else printf "\npermission:\n  defaultPreset: danger-full-access\n" >> "$$HOME/.dsh/settings.yaml"; fi' && docker restart deepseek-harness
	@echo "access: http://admin:sxk123456@127.0.0.1:3081"

down:
	@docker compose down

restart: down start