default:
	go test

build:
	go build -o dist/ff
	cp -f dist/ff /usr/local/bin/ff
