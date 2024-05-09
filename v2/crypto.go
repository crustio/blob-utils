package v2

import (
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"
)

func EncodeBlobs(data []byte) []kzg4844.Blob {
	blobs := []kzg4844.Blob{{}}
	blobIndex := 0
	fieldIndex := -1
	for i := 0; i < len(data); i += 31 {
		fieldIndex++
		if fieldIndex == params.BlobTxFieldElementsPerBlob {
			blobs = append(blobs, kzg4844.Blob{})
			blobIndex++
			fieldIndex = 0
		}
		max := i + 31
		if max > len(data) {
			max = len(data)
		}
		copy(blobs[blobIndex][fieldIndex*32:], data[i:max])
	}
	return blobs
}

func DecodeBlob(blob []byte) []byte {
	if len(blob) != params.BlobTxFieldElementsPerBlob*32 {
		panic("invalid blob encoding")
	}
	var data []byte

	// XXX: the following removes trailing 0s in each field element (see EncodeBlobs), which could be unexpected for certain blobs
	j := 0
	for i := 0; i < params.BlobTxFieldElementsPerBlob; i++ {
		data = append(data, blob[j:j+31]...)
		j += 32
	}

	i := len(data) - 1
	for ; i >= 0; i-- {
		if data[i] != 0x00 {
			break
		}
	}
	data = data[:i+1]
	return data
}
