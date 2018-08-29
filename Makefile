current_dir=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
git_tag=$(shell git describe --tags --abbrev^=0 2> /dev/null)
git_hash=$(shell git rev-parse HEAD 2> /dev/null)
time=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

PACKAGE:=$(subst ${GOPATH}/src/,,${current_dir})
PROJECT_NAME:=$(notdir ${current_dir})
REGISTRY_PATH:=

ifeq (${git_tag}, )
VERSION:=0.0
else
VERSION:=${git_tag}
endif

.PHONY: build
pre-build:
	@glide install
build: pre-build
	@docker build \
		--build-arg projectPath=${PACKAGE} \
		--build-arg version=${VERSION} \
		--build-arg commit=${git_hash} \
		--build-arg buildTime=${time} \
		-t "${REGISTRY_PATH}${PROJECT_NAME}:${VERSION}" \
		.
	@docker tag "${REGISTRY_PATH}${PROJECT_NAME}:${VERSION}" "${REGISTRY_PATH}${PROJECT_NAME}:latest"