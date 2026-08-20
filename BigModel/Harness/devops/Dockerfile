# syntax=docker/dockerfile:1.7

ARG BASE_IMAGE=runzhliu/deepseek-harness:0.1.0-rc.6

FROM ${BASE_IMAGE}

USER root

ARG TARGETARCH
ARG GO_VERSION=latest
ARG GO_REFRESH=unknown
ARG BUILD_DATE=unknown
ARG VCS_REF=unknown

LABEL org.opencontainers.image.title="DeepSeek Harness Dev Base" \
      org.opencontainers.image.description="DeepSeek Harness development environment with Go, Python, C/C++ and Docker CLI" \
      org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.revision="${VCS_REF}"


# ============================================================
# Basic Development / System Tools
# ============================================================
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y --no-install-recommends \
        \
        ca-certificates \
        curl \
        wget \
        \
        git \
        git-lfs \
        openssh-client \
        \
        jq \
        vim \
        nano \
        less \
        tree \
        file \
        \
        unzip \
        zip \
        tar \
        rsync \
        \
        ripgrep \
        \
        procps \
        htop \
        lsof \
        strace \
        \
        net-tools \
        iputils-ping \
        dnsutils \
    && rm -rf /var/lib/apt/lists/*


# ============================================================
# Python Development Environment
# ============================================================
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y --no-install-recommends \
        python3 \
        python3-dev \
        python3-pip \
        python3-venv \
        python-is-python3 \
    && rm -rf /var/lib/apt/lists/*


# ============================================================
# C / C++ Development Environment
# ============================================================
#
# gcc / g++
# clang / clangd
# gdb / lldb
# cmake / ninja / meson
# clang-format / clang-tidy
# valgrind / cppcheck
# autoconf / automake
#
# ============================================================
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y --no-install-recommends \
        \
        build-essential \
        gcc \
        g++ \
        make \
        \
        gdb \
        \
        clang \
        clangd \
        clang-format \
        clang-tidy \
        lldb \
        \
        cmake \
        ninja-build \
        meson \
        \
        pkg-config \
        autoconf \
        automake \
        libtool \
        \
        ccache \
        cppcheck \
        valgrind \
        \
        libssl-dev \
        zlib1g-dev \
    && rm -rf /var/lib/apt/lists/*


# ============================================================
# Docker CLI / Docker Compose / Docker Buildx
# ============================================================
#
# 这里只安装 Docker Client。
#
# 不在 Harness 容器中启动 dockerd。
#
# Runtime 时通过：
#
#   /var/run/docker.sock
#
# 操作宿主机 Docker Engine。
#
# ============================================================
RUN set -eux; \
    . /etc/os-release; \
    \
    case "${ID}" in \
        debian|ubuntu) ;; \
        *) \
            echo "Unsupported Linux distribution: ${ID}"; \
            exit 1; \
            ;; \
    esac; \
    \
    install -m 0755 -d /etc/apt/keyrings; \
    \
    curl -fsSL "https://download.docker.com/linux/${ID}/gpg" \
        -o /etc/apt/keyrings/docker.asc; \
    \
    chmod a+r /etc/apt/keyrings/docker.asc; \
    \
    echo \
        "deb [arch=${TARGETARCH} signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/${ID} ${VERSION_CODENAME} stable" \
        > /etc/apt/sources.list.d/docker.list; \
    \
    apt-get update; \
    \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y --no-install-recommends \
        docker-ce-cli \
        docker-buildx-plugin \
        docker-compose-plugin; \
    \
    rm -rf /var/lib/apt/lists/*


# ============================================================
# Latest Stable Go
# ============================================================
#
# 默认：
#
#   GO_VERSION=latest
#
# 自动从：
#
#   https://go.dev/VERSION?m=text
#
# 获取官方最新 stable。
#
# 也可以指定：
#
#   --build-arg GO_VERSION=go1.xx.x
#
# TARGETARCH 由 Docker Buildx 自动提供：
#
#   amd64
#   arm64
#
# GO_REFRESH 用于打破 Docker cache，
# Makefile 默认每天变化一次。
#
# ============================================================
RUN set -eux; \
    \
    echo "GO_REFRESH=${GO_REFRESH}"; \
    echo "TARGETARCH=${TARGETARCH}"; \
    \
    case "${TARGETARCH}" in \
        amd64|arm64) ;; \
        *) \
            echo "Unsupported Go architecture: ${TARGETARCH}"; \
            exit 1; \
            ;; \
    esac; \
    \
    if [ "${GO_VERSION}" = "latest" ]; then \
        RESOLVED_GO_VERSION="$(curl -fsSL 'https://go.dev/VERSION?m=text' | head -n 1)"; \
    else \
        RESOLVED_GO_VERSION="${GO_VERSION}"; \
    fi; \
    \
    echo "Installing ${RESOLVED_GO_VERSION} linux/${TARGETARCH}"; \
    \
    GO_FILE="${RESOLVED_GO_VERSION}.linux-${TARGETARCH}.tar.gz"; \
    \
    curl -fsSL \
        "https://go.dev/dl/${GO_FILE}" \
        -o /tmp/go.tar.gz; \
    \
    GO_SHA256="$( \
        curl -fsSL 'https://go.dev/dl/?mode=json&include=all' \
        | jq -r \
            --arg version "${RESOLVED_GO_VERSION}" \
            --arg arch "${TARGETARCH}" \
            '.[] \
            | select(.version == $version) \
            | .files[] \
            | select( \
                .os == "linux" \
                and .arch == $arch \
                and .kind == "archive" \
            ) \
            | .sha256' \
        | head -n 1 \
    )"; \
    \
    test -n "${GO_SHA256}"; \
    test "${GO_SHA256}" != "null"; \
    \
    echo "${GO_SHA256}  /tmp/go.tar.gz" \
        | sha256sum -c -; \
    \
    rm -rf /usr/local/go; \
    \
    tar \
        -C /usr/local \
        -xzf /tmp/go.tar.gz; \
    \
    rm -f /tmp/go.tar.gz


# ============================================================
# Go Environment
# ============================================================

ENV GOPATH=/opt/go
ENV GOBIN=/usr/local/bin
ENV GOCACHE=/var/cache/go-build
ENV GOMODCACHE=/opt/go/pkg/mod

ENV PATH="/usr/local/go/bin:/usr/local/bin:${PATH}"

RUN mkdir -p \
        "${GOPATH}" \
        "${GOMODCACHE}" \
        "${GOCACHE}" \
    && chmod -R 0777 \
        "${GOPATH}" \
        "${GOCACHE}"


# ============================================================
# Go Development Tools
# ============================================================
#
# gopls       Language Server
# goimports   import + format
# staticcheck 静态检查
# dlv         Debugger
#
# ============================================================
RUN set -eux; \
    go version; \
    \
    go install golang.org/x/tools/gopls@latest; \
    go install golang.org/x/tools/cmd/goimports@latest; \
    go install honnef.co/go/tools/cmd/staticcheck@latest; \
    go install github.com/go-delve/delve/cmd/dlv@latest


# ============================================================
# Workspace
# ============================================================

RUN mkdir -p /workspace

WORKDIR /workspace


# ============================================================
# Build-time Smoke Test
# ============================================================
#
# 确保基础开发环境在 image build 阶段就是完整的。
#
# ============================================================
RUN set -eux; \
    \
    echo "============================"; \
    echo "Git"; \
    echo "============================"; \
    git --version; \
    \
    echo "============================"; \
    echo "Python"; \
    echo "============================"; \
    python --version; \
    pip3 --version; \
    \
    echo "============================"; \
    echo "Go"; \
    echo "============================"; \
    go version; \
    gopls version; \
    staticcheck -version; \
    dlv version; \
    \
    echo "============================"; \
    echo "C / C++"; \
    echo "============================"; \
    gcc --version | head -n 1; \
    g++ --version | head -n 1; \
    clang --version | head -n 1; \
    cmake --version | head -n 1; \
    ninja --version; \
    \
    echo "============================"; \
    echo "Docker"; \
    echo "============================"; \
    docker --version; \
    docker compose version; \
    docker buildx version


# ============================================================
# Runtime
# ============================================================
#
# Harness 作为高权限开发 Agent 使用。
#
# 保留基础镜像原有 ENTRYPOINT/CMD。
#
# ============================================================

USER root

WORKDIR /workspace