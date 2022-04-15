.PHONY: run build docker docker-release

#当前目录
CURDIR := $(shell pwd)
# 入口 main.go 所在文件夹路径
MAIN_PATH := $(if $(MAIN_PATH),$(MAIN_PATH),$(CURDIR))
# 可执行文件名称
APP_NAME := $(if $(APP_NAME),$(APP_NAME),app)
# 服务名称
PROJECT_NAME := $(shell echo $(CURDIR)|awk -F '/' '{ print $$NF }')
# 提交版本
COMMIT_ID := $(shell git rev-parse --short HEAD)
# 打包时间
BUILD_DATE := $(shell date "+%Y-%m-%d/%H:%M:%S")
# 平台架构
ARCH := $(if $(ARCH),$(ARCH),linux/amd64)
# 编译时参数
LDFLAGS := "-w -s -extldflags -static -X github.com/gotomicro/ego/core/eapp.appName='${APP_NAME}' -X github.com/gotomicro/ego/core/eapp.buildVersion='${COMMIT_ID}' -X github.com/gotomicro/ego/core/eapp.buildAppVersion='${COMMIT_ID}' -X github.com/gotomicro/ego/core/eapp.buildStatus='Modified' -X github.com/gotomicro/ego/core/eapp.buildTag='${TARGET}' -X github.com/gotomicro/ego/core/eapp.buildUser='doraemon' -X github.com/gotomicro/ego/core/eapp.buildHost='127.0.0.1' -X github.com/gotomicro/ego/core/eapp.buildTime='${BUILD_DATE}'"

build:
	@GOARCH=$(shell echo $(ARCH) | cut -d "/" -f 2) go build -mod vendor -ldflag $(LDFLAGS) -o $(CURDIR)/$(PROJECT_NAME) $(MAIN_PATH)/main.go

docker:
