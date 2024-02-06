# Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# Build all by default, even if it's not first
.DEFAULT_GOAL := all

ROOT_PACKAGE=github.com/JoyZF/zoom
VERSION_PACKAGE=github.com/marmotedu/component-base/pkg/version

.PHONY: all
# TODO lint cover build
all: add-copyright format lint cover build

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/copyright.mk
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk


## add-copyright: Ensures source code files have copyright license headers.
.PHONY: add-copyright
add-copyright:
	@$(MAKE) copyright.add


## format: Gofmt (reformat) package sources (exclude vendor dir if existed).
.PHONY: format
format: tools.verify.golines tools.verify.goimports
	@echo "===========> Formating codes"
	@$(FIND) -type f -name '*.go' | $(XARGS) gofmt -s -w
	@$(FIND) -type f -name '*.go' | $(XARGS) goimports -w -local $(ROOT_PACKAGE)
	@$(FIND) -type f -name '*.go' | $(XARGS) golines -w --max-len=120 --reformat-tags --shorten-comments --ignore-generated .
	@$(GO) mod edit -fmt


## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	echo $(MAKE)
	@$(MAKE) go.lint

## cover: Run unit test and get test coverage.
.PHONY: cover
cover:
	@$(MAKE) go.test.cover

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build


## build: Run swag-fmt
.PHONY: swag-fmt
swag-fmt:
	swag fmt -g cmd/zoom-apiserver/main.go

## build: Run swag-init
.PHONY: swag-init
swag-init:
	swag init -g cmd/zoom-apiserver/main.go

.PHONY: swag
swag: swag-fmt swag-init
