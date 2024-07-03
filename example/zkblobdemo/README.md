## zkblob demo


build: ```make build```

usage:

1. get proof:

./zkblobdemo [blob-hash]

```shell
./zkblobdemo 0x014a394ce48abea40e57a780ce08a0f58789b2a966be874c85d3ae466858a8bb
```

2. verify:

./zkblobdemo [batch-num] [blob-hash] [proof-list]

```shell
./zkblobdemo 508 0x014a394ce48abea40e57a780ce08a0f58789b2a966be874c85d3ae466858a8bb 0x014a394ce48abea40e57a780ce08a0f58789b2a966be874c85d3ae466858a8bb
```