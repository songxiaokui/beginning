.PHONY: build-arm64 build-amd64 download_arm64

build-arm64:
	@echo "build start time: $$(date '+%Y-%m-%d %H:%M:%S')"
	@singularity build \
         --build-arg target_arch=arm64 \
         --build-arg build_jobs=4 \
         openfoam-13-arm64.sif \
         openfoam13-multiarch.def
	@echo "build end time: $$(date '+%Y-%m-%d %H:%M:%S')"

build-amd64:
	@echo "build start time: $$(date '+%Y-%m-%d %H:%M:%S')"
	@eval 'HTTP_PROXY="http://192.168.22.50:8889" HTTPS_PROXY="http://192.168.22.50:8889"' singularity build \
         --build-arg target_arch=amd64 \
         --build-arg build_jobs=4 \
         openfoam-13-amd64.sif \
         openfoam13-multiarch.def
	@echo "build end time: $$(date '+%Y-%m-%d %H:%M:%S')"

download_arm64:
	@scp -r qls:/home/cc/Desktop/qy_parallel/openfoam/openfoam_13_arrch64.sif ./images/