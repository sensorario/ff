default:
	go test ./... -v

build:
	env GO111MODULE=on go build -o ff ./...
	sudo mv ff /usr/local/bin/ff
	cd /usr/local/bin/ && ln -sf /usr/local/bin/ff git-ff

update:
	go get -u ./...
