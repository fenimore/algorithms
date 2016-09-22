package convert

import "testing"

func TestHexToBase64(t *testing.T) {
	given := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result, err := HexToBase64OfString(given)
	if err != nil {
		t.Error(err)
	}
	if string(result) != expected {

	}

	result, err = HexToBase64OfByte([]byte(given))
	if err != nil {
		t.Error(err)
	}
	if string(result) != expected {
		t.Error("Unexecpted %s", string(result))
	}
}
