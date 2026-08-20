.PHONY: download build-kylin-arm64 build-amd64 help

SIF_DEF := openfoam-v2606-multiarch.def
BUILD_JOBS ?= 4
LOG_DIR ?= logs

download:  ## 下载 OpenCFD OpenFOAM v2606 源码
	@mkdir -p tar
	@test -f tar/OpenFOAM-v2606.tgz || wget -O tar/OpenFOAM-v2606.tgz https://sourceforge.net/projects/openfoam/files/v2606/OpenFOAM-v2606.tgz/download
	@test -f tar/ThirdParty-v2606.tgz || wget -O tar/ThirdParty-v2606.tgz https://netactuate.dl.sourceforge.net/project/openfoam/v2606/ThirdParty-v2606.tar.gz

build-kylin-arm64: download  ## 在麒麟 ARM64 主机上构建
	@mkdir -p $(LOG_DIR)
	@set +e; singularity build --build-arg target_arch=arm64 --build-arg build_jobs=$(BUILD_JOBS) openfoam-v2606-kylin-arm64.sif $(SIF_DEF) > $(LOG_DIR)/openfoam-v2606-kylin-arm64.stdout.log 2> $(LOG_DIR)/openfoam-v2606-kylin-arm64.stderr.log; status=$$?; \
	if [ $$status -ne 0 ]; then echo "build failed: see $(LOG_DIR)/openfoam-v2606-kylin-arm64.stdout.log and $(LOG_DIR)/openfoam-v2606-kylin-arm64.stderr.log"; fi; exit $$status

build-amd64: download  ## 在 AMD64 主机上构建
	@mkdir -p $(LOG_DIR)
	@set +e; eval 'HTTP_PROXY="http://192.168.22.50:8889" HTTPS_PROXY="http://192.168.22.50:8889"' singularity build --build-arg target_arch=amd64 --build-arg build_jobs=$(BUILD_JOBS) openfoam-v2606-amd64.sif $(SIF_DEF) > $(LOG_DIR)/openfoam-v2606-amd64.stdout.log 2> $(LOG_DIR)/openfoam-v2606-amd64.stderr.log; status=$$?; \
	if [ $$status -ne 0 ]; then echo "build failed: see $(LOG_DIR)/openfoam-v2606-amd64.stdout.log and $(LOG_DIR)/openfoam-v2606-amd64.stderr.log"; fi; exit $$status

help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ { printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
