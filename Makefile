.PHONY: build run test
PACKAGES=$(shell GO111MODULE=on go list -mod=vendor ./... | grep -v '/vendor/')

build:
	@echo Version: ${VERSION}
	@GO111MODULE=on go build cmd/*

run:
	@GO111MODULE=on go run -mod=vendor -v cmd/main.go

test:
	@GO111MODULE=on go test -mod=vendor ${PACKAGES} && \
	echo "\033[1;92m OK \033[0m" || (echo "\033[1;91m FAIL \033[0m" && exit 1)

generate_docs:
	swag init -g service/api.go && \
	rm docs/swagger.json docs/docs.go

generate_mocks:
	cd service/ && mockery --all