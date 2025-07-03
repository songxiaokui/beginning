
VERSION ?= 'v5'
FQ ?= 1

.PHONY: run
run:  ## 运行测试案例
	@echo "运行测试案例"
	@#. /etc/profile && matlab -batch "addpath('$(CURDIR)'); generate_astro_data('2025-05-23 00:00:00','2025-05-23 00:00:02',2000,'./testcase/output/tes.mat', '${VERSION}', ${FQ})"
	@. /etc/profile && matlab -batch "addpath('$(CURDIR)'); generate_astro_data()"

.PHONY: main
main:
	@echo "Build and Running"
	@. /etc/profile && matlab -nodisplay -nosplash -batch "addpath('$(CURDIR)'); cmdx;"


help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
		printf "\033[36m%-30s\033[0m \033[31m%s\033[0m\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help