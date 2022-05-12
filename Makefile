COMPILE_TIME=$(shell date "+%Y-%m-%d")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null)
BUILT_TYPE="develop"
Version="v1.0.0"
ldflags="-X github.com/HiData-xyz/go-example/build.Commit=$(COMMIT) -X github.com/HiData-xyz/go-example/build.BuildType=$(BUILT_TYPE) -X github.com/HiData-xyz/go-example/build.V=$(Version)"

.PHONY: build
build:
	GOOS=linux go build -ldflags $(ldflags) -o HiData.Manager

docker: build
	docker build . -t hidata/manager:v2.0.0
	docker image save -o "hidata-manager_$(COMPILE_TIME).tar.gz" "hidata/manager:v2.0.0"
