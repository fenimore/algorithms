package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// PKSPad takes in plaintext and returns the padded text.
func PKSPad(plaintext []byte, blocksize int) []byte {
	remainder := len(plaintext) % blocksize
	if remainder == 0 {
		return plaintext
	}

	// how much to pad
	padSize := blocksize - remainder
	padding := make([]byte, padSize)
	for i := range padding {
		padding[i] = byte(padSize)
	}

	return append(plaintext, padding...)
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

	for i := 0; i*size < len(plaintext); i++ {
		cipher.Encrypt(ciphertext[i*size:size*(i+1)], plaintext[i*size:size*(i+1)])
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

func EncryptCBC(key, plaintext, iv []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func RandomKey(size int) []byte {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = byte(r.Intn(256))
	}

	return result
}

func BlackBox(plaintext []byte) []byte {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	// random padding
	bufSize := r.Intn(11)
	if bufSize < 5 {
		bufSize += 5
	}
	buf := make([]byte, bufSize)
	// padd and then pad again
	plaintext = append(plaintext, buf...)
	plaintext = append(buf, plaintext...)
	plaintext = PKSPad(plaintext, 16)
	// generate random keys etcy
	key := RandomKey(16)
	iv := RandomKey(16)

	which := r.Intn(2)
	if which == 1 {
		fmt.Println("CBC encrypting")
		return EncryptCBC(key, plaintext, iv)
	}
	fmt.Println("ECB encrypting")
	return EncryptECB(key, plaintext)
}

func DetectECB(text []byte) bool {
	amt := len(text) / 16 // blocksize is going to be 16
	repeats := make(map[string]bool)
	for i := 0; i < amt; i++ {
		if repeats[string(text[i*16:i*16+16])] {
			return true
		}
		repeats[string(text[i*16:i*16+16])] = true
	}
	return false
}

func Oracle(plaintext []byte) []byte {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	// random padding
	bufSize := r.Intn(11)
	if bufSize < 5 {
		bufSize += 5
	}
	buf := make([]byte, bufSize)
	// padd and then pad again
	plaintext = append(plaintext, buf...)
	plaintext = append(buf, plaintext...)
	plaintext = PKSPad(plaintext, 16)
	// generate random keys etcy
	key := globalkey

	return EncryptECB(key, plaintext)
}

var (
	file     *os.File
	key      []byte
	result   []byte
	expected []byte
	err      error

	plaintext  []byte
	ciphertext []byte

	globalkey []byte
)

func main() {
	// challenge 9
	fmt.Println(string(PKSPad([]byte("YELLOW SUBMARINE"), 20)))
	fmt.Println(PKSPad([]byte("YELLOW SUBMARINE"), 20))
	// challenge 10
	fmt.Println("Challenge 10")
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
	fmt.Println(ciphertext[:10])
	// initialization vector is all ASCII zeroes
	iv := make([]byte, len(key))
	plaintext := DecryptCBC(key, ciphertext, iv)
	fmt.Println(string(plaintext[:15])) // Play that funky music white boy
	ciphertext = EncryptCBC(key, plaintext, iv)
	fmt.Println(ciphertext[:10])
	plaintext = DecryptCBC(key, ciphertext, iv)
	fmt.Println(string(plaintext[:15]))
	// Challenge 11
	fmt.Println("Challenge 11")
	fmt.Println(DetectECB(BlackBox(plaintext)))
	// challenge 12 byte at a time
	fmt.Println("Challenge 12")
	// 64 len []byte:
	plaintext = []byte("twas brillig and the slithey toves, did gyre and gimble in the w")
	globalkey = RandomKey(16)
	str := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg"
	str += "aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq"
	str += "dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg"
	str += "YnkK"

	padbyte, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(err)
	}

	encryptedAlone := Oracle(plaintext)
	encryptedTogether := Oracle(append(plaintext, padbyte...))
	fmt.Println(len(encryptedTogether), len(encryptedAlone))
}
