## Download
```shell
docker pull --platform=linux/arm64 ghcr.io/ollama-webui/ollama-webui:main
docker save -o ollama-webui-arm64.tar ghcr.io/ollama-webui/ollama-webui:main
docker inspect ghcr.io/ollama-webui/ollama-webui:main | grep Architecture
```