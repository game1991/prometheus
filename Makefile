.PHONY: run build docker docker-release

IMAGE :=golang:1.18-buster
# 入口 main.go 所在文件夹路径
MAIN_PATH := $(if $(MAIN_PATH),$(MAIN_PATH),$(CURDIR))
# 可执行文件名称
APP_NAME := $(if $(APP_NAME),$(APP_NAME),app)
# 服务名称
PROJECT_NAME := $(shell echo $(shell pwd)|awk -F '/' '{ print $$NF }')
# 提交版本
COMMIT_ID := $(shell git log -n 1 --oneline)
# 网卡信息
ETHERNET := $(if $(ETHERNET),$(ETHERNET),eno1)
# 平台架构
ARCH := $(if $(ARCH),$(ARCH),$(shell uname -r))
# 打包主机
BUILD_HOST := $(shell ifconfig $(ETHERNET) | grep -w 'inet' | awk '{print $$2}')
# 打包时间
BUILD_DATE := $(shell date "+%Y-%m-%d/%H:%M:%S")
# 打包者
BUILD_USER := $(if $(BUILD_USER),$(BUILD_USER),localuser)
# 镜像地址
DOCKER_BASE:=$(if $(DOCKER_BASE),$(DOCKER_BASE),local/app)
# 镜像便签
DOCKER_TAG:=$(if $(DOCKER_TAG),$(DOCKER_TAG),$(shell echo $(ARCH)|awk '{ sub(/\//,"-"); print $$0 }'))
# 构建插件
DOCKER_BUILDKIT:=$(if $(DOCKER_BUILDKIT),$(DOCKER_BUILDKIT),DOCKER_BUILDKIT=1)

# 编译时参数
LDFLAGS :="-w -s -extldflags '-static' -X 'github.com/gotomicro/ego/core/eapp.appName=${APP_NAME}' -X 'github.com/gotomicro/ego/core/eapp.buildAppVersion=${COMMIT_ID}' -X 'github.com/gotomicro/ego/core/eapp.buildStatus=Modified' -X 'github.com/gotomicro/ego/core/eapp.buildUser=${BUILD_USER}' -X 'github.com/gotomicro/ego/core/eapp.buildHost=${BUILD_HOST}' -X 'github.com/gotomicro/ego/core/eapp.buildTime=${BUILD_DATE}'"

build:
	go build -mod vendor -ldflags ${LDFLAGS} -o $(PROJECT_NAME) main.go

docker:
	$(DOCKER_BUILDKIT) docker build \
	--no-cache \
	--build-arg IMAGE=$(IMAGE) \
	--build-arg APP_NAME=$(APP_NAME) \
	--build-arg PROJECT_NAME=$(PROJECT_NAME) \
	-t $(DOCKER_BASE):$(DOCKER_TAG) \
	-f Dockerfile . || exit 1