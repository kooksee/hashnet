.PHONY: version build build_linux docker_login docker_build docker_push_dev docker_push_pro
.PHONY: rm_stop

_Version = v1.0.0
_VersionFile = version/version.go
_CommitVersion = $(shell git rev-parse --short=8 HEAD)
_BuildVersion = $(shell date "+%F %T")
_GOBIN = $(shell pwd)

_ImageName = hashnet
_ProjectPath = github.com/kooksee/$(_ImageName)
_ImagesPrefix = registry.cn-hangzhou.aliyuncs.com/ybase/
_ImageLatestName = $(_ImagesPrefix)$(_ImageName)
_ImageTestName = $(_ImagesPrefix)$(_ImageName):test
_ImageVersionName = $(_ImagesPrefix)$(_ImageName):$(_Version)

_version:
	@echo "项目版本处理"
	@echo "package version" > $(_VersionFile)
	@echo "const Version = "\"$(_Version)\" >> $(_VersionFile)
	@echo "const BuildVersion = "\"$(_BuildVersion)\" >> $(_VersionFile)
	@echo "const CommitVersion = "\"$(_CommitVersion)\" >> $(_VersionFile)

b:
	@echo "开始编译"
	GOBIN=$(_GOBIN) go install main.go

_build_linux: _version
	@echo "交叉编译成linux应用"
	docker run --rm -v $(GOPATH):/go golang:latest go build -o /go/src/$(_ProjectPath)/main /go/src/$(_ProjectPath)/main.go

rm_stop:
	@echo "删除所有的的容器"
	sudo docker rm -f $(sudo docker ps -qa)
	sudo docker ps -a

rm_none:
	@echo "删除所为none的image"
	sudo docker images  | grep none | awk '{print $3}' | xargs docker rmi -f

docker:_build_linux
	@echo "docker build and push"
	sudo docker build -t $(_ImageVersionName) .
	sudo docker push $(_ImageVersionName)
