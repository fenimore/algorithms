package convert

import (
	"encoding/hex"
	"errors"
)

func FixedXOR(a, b []byte) ([]byte, error) {
	first := make([]byte, hex.DecodedLen(len(a)))
	second := make([]byte, hex.DecodedLen(len(b)))
	count1, _ := hex.Decode(first, a)
	count2, _ := hex.Decode(second, b)
	if count1 != count2 {
		return nil, errors.New("Buffers not the same size")
	}
	result := make([]byte, count1)
	for i := range result {
		result[i] = first[i] ^ second[i]
	}
	return result, nil
}
