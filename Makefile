TEST?=$$(go list ./... | grep -v 'vendor')
TESTNAME?=
TF_LOG?=INFO
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
BUDDY_TOKEN?=1234567890
BUDDY_BASE_URL?=https://api.buddy.works
BUDDY_INSECURE?=false
BUDDY_GET_TOKEN?=curl
BUDDY_GH_TOKEN?=ABCDEFGH
BUDDY_GH_PROJECT?=test/test

default: lint

test_dev:
	$(eval BUDDY_TOKEN=$(shell sh -c "${BUDDY_GET_TOKEN}"))
	go clean -testcache
	TF_LOG=${TF_LOG} BUDDY_TOKEN=${BUDDY_TOKEN} BUDDY_GH_TOKEN=${BUDDY_GH_TOKEN} BUDDY_GH_PROJECT=${BUDDY_GH_PROJECT} BUDDY_BASE_URL=https://api.awsdev.net BUDDY_INSECURE=true go test $(TEST) -v ${TESTNAME} -timeout 60m

test:
	go clean -testcache
	TF_LOG=${TF_LOG} BUDDY_TOKEN=${BUDDY_TOKEN} BUDDY_GH_TOKEN=${BUDDY_GH_TOKEN} BUDDY_GH_PROJECT=${BUDDY_GH_PROJECT} BUDDY_BASE_URL=${BUDDY_BASE_URL} BUDDY_INSECURE=${BUDDY_INSECURE} go test $(TEST) -v ${TESTNAME} -timeout 60m

fmt:
	gofmt -w $(GOFMT_FILES)

lint: fmt golangci

golangci:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run ./...

.PHONY: default build test_dev test fmt lint golangci