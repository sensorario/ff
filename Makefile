default:
	go test

build:
	sudo env GO111MODULE=on go build -o /usr/local/bin/ff
