.PHONY: build

override TARGETOS := linux
override TARGETARCH := amd64

build:
	@GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o cmd_${TARGETOS}-${TARGETARCH} cmd.go
	@chmod +x cmd_$TARGETOS-$TARGETARCH
	@scp -r cmd_${TARGETOS}-${TARGETARCH} simcae-prod:/home/cc/TestData/
	@scp -r config.yaml simcae-prod:/home/cc/TestData/
	@echo "run use: cd /home/cc/TestData; ./cmd_${TARGETOS}-${TARGETARCH} --config config.yaml"

clean:
	@rm -f cmd_$TARGETOS-$TARGETARCH