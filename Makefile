default: build
	./v2/exa tx-simple

build:
	go build -ldflags="-w -s" -o v2/exa github.com/crustio/blob-utils/v2/cmd
