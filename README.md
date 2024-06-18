# Blob Utils

CLI for messing around with Proto-Danksharding blobs. To get started run:
```
git clone https://github.com/crustio/blob-utils && cd blob-utils
go build
./blob-utils -help
```

## Features
- Creating and sending blob transactions
- Download blobs sidecars

Feel free to open an issue request for more features.

(thanks to @mdehoog for the initial implementation)

## Blob Example

> Note: Before running the Example, you may need to modify the `BlobAddress` at the beginning of `blob.go`

- before run

modify rpcURL and private key in the source code `example/main.go`

- run example:

```shell
cd example && go run main.go
```