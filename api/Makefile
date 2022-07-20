# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

RUNTIME_DIR ?= .

build: go

go:
	@find . -name '*.pb.go' -delete
	docker run -it \
		-v `pwd`:/build \
		-v $(RUNTIME_DIR)/api:/api \
		atomix/codegen:go-latest \
	    --proto-path . --proto-path /api --proto-files atomix/gossip/v1/*.proto --go-path . --import-path github.com/atomix/gossip/api
