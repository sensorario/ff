default:
	go test ./... -v

update_mod:
	GO111MODULE=on go mod tidy

install:
	env GO111MODULE=on go build -o ff ./src/*
	sudo mv ff /usr/local/bin/ff

update:
	go get -u ./...
