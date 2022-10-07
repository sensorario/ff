default:
	go test ./... -v

update_mod:
	GO111MODULE=on go mod tidy

install:
	env GO111MODULE=on go build -o ff ./cmd/ff/*.go
	sudo mv ff /usr/local/bin/ff

update:
	go get -u ./...

GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test ./...
.PHONY: test/cover
test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out
