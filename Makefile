default:
	go test

build:
	env GO111MODULE=on go build -o dist/ff
	cp -f dist/ff /usr/local/bin/ff
