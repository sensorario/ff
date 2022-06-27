default:
	go test ./... -v

install:
	env GO111MODULE=on go build -o ff ./...
	sudo mv ff /usr/local/bin/ff

update:
	go get -u ./...
