default:
	go test

build:
	env GO111MODULE=on go build -o ff ./...
	sudo mv ff /usr/local/bin/ff
