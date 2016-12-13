package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func PadPKCS(block []byte, length int) []byte {
	buffer := make([]byte, length)
	if len(block) == length {
		return block
	}
	copy(buffer[:len(block)], block)
	// set last byte
	lastByte := byte(length - len(block))
	buffer[len(buffer)-1] = lastByte

	return buffer
}

func DecryptECB(key, ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	cipher, err := aes.NewCipher(key)
	size := cipher.BlockSize()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i*size < len(ciphertext); i++ {
		cipher.Decrypt(plaintext[i*size:size*(i+1)], ciphertext[i*size:size*(i+1)])
		//fmt.Println(string(plaintext[i*size : size*(i+1)]))
	}
	return plaintext
}

func EncryptECB(key, plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	size := cipher.BlockSize()

	fmt.Println(len(ciphertext), len(plaintext))
	for i := 0; i*size < len(plaintext); i++ {
		cipher.Encrypt(ciphertext[i*size:size*(i+1)], plaintext[i*size:size*(i+1)])
		//fmt.Println(ciphertext[i*size:size*i+1], plaintext[i*size:size*i+1])
	}

	return ciphertext
}

func DecryptCBC(key, ciphertext, iv []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(plaintext, ciphertext)
	return plaintext
}

// func EncryptCBC(key, plaintext []byte) []byte {
//	ciphertext := make([]byte, len(plaintext))
//	return ciphertext
// }

var (
	file     *os.File
	key      []byte
	result   []byte
	expected []byte
	err      error

	plaintext  []byte
	ciphertext []byte
)

func main() {
	// challenge 9
	fmt.Println(PadPKCS([]byte("YELLOW SUBMARINE"), 20))
	// challenge 10
	key = []byte("YELLOW SUBMARINE")
	file, err = os.Open("inputs/challenge_10.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	ciphertext, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		fmt.Println(err)
	}
	// initialization vector is all ASCII zeroes
	iv := make([]byte, len(key))
	plaintext := DecryptCBC(key, ciphertext, iv)
	fmt.Println(plaintext[:10]) // Play that funky music white boy

	// Challenge 11

}
