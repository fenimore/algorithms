package convert

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(input string) string {
	var data []byte
	var result []byte
	data, err := hex.DecodeString(input)
	base64.StdEncoding.Encode(result, data)
	return result
}
