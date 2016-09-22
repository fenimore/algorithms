package bytes

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// HexToBase64 takes in a hex and converts it to base64.
func HexToBase64(input []byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(input)))

	_, err := hex.Decode(data, input)
	if err != nil {
		return nil, err
	}
	result := make([]byte, base64.StdEncoding.EncodedLen(len(data)))

	base64.StdEncoding.Encode(result, data)
	return result, nil
}

// StringHexToBase64 wrapper around HexToBase64 for using
// strings instead of byte buffers.
func StringHexToBase64(input string) (string, error) {
	result, err := HexToBase64([]byte(input))
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// FixedXORHex takes two hex of equal length
// and returns the hex result of an XOR operation.
func FixedXORHex(a, b []byte) ([]byte, error) {
	first := make([]byte, hex.DecodedLen(len(a)))
	second := make([]byte, hex.DecodedLen(len(b)))
	count1, _ := hex.Decode(first, a)
	count2, _ := hex.Decode(second, b)
	if count1 != count2 {
		return nil, errors.New("Buffers not the same size")
	} else if count1 < 1 {
		return nil, errors.New("Buffer zero length")
	}
	data := make([]byte, count1)
	for i := range data {
		data[i] = first[i] ^ second[i]
	}
	result := make([]byte, hex.EncodedLen(len(data)))
	_ = hex.Encode(result, data)

	return result, nil
}

// Check a hex against a XOR cipher of a single character.
func SingleByteXORCipher(h []byte, cipher byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(h)))
	_, err := hex.Decode(data, h)
	if err != nil {
		return nil, err
	}
	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ cipher
	}

	return result, nil
}
