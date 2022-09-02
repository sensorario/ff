default:
	go test ./... -v

update_mod:
	GO111MODULE=on go mod tidy

install: update_mod
	env GO111MODULE=on go build -o ff ./...
	sudo mv ff /usr/local/bin/ff

update: update_mod
	go get -u ./...
