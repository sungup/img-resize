install:
  @go install github.com/sungup/img-resize/cmd/img-resize

build:
  @go build -o bin/img-resize github.com/sungup/img-resize/cmd/img-resize

vet:
  @go vet $$(go list ./... | grep -v img-resize)

vendor:
  @go mod vendor

.PHONY: vendor install vet