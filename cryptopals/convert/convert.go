package convert

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64OfString(input string) ([]byte, error) {
	var data []byte
	data, err := hex.DecodeString(input)
	if err != nil {
		return nil, err
	}
	result := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(result, data)

	return result, nil
}

func HexToBase64OfByte(input []byte) ([]byte, error) {
	data := make([]byte, hex.DecodedLen(len(input)))

	_, err := hex.Decode(data, input)
	if err != nil {
		return nil, err
	}
	result := make([]byte, base64.StdEncoding.EncodedLen(len(data)))

	base64.StdEncoding.Encode(result, data)
	return result, nil
}
