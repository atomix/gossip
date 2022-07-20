# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

build:
	go build ./...

generate: go

go:
	@find . -name '*.pb.go' -delete
	docker run -it \
		-v `pwd`:/build \
		atomix/codegen:go-latest \
	    --proto-path . --go-path . --import-path github.com/atomix/gossip-storage/api
