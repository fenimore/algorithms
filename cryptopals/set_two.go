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
	for i := 0; i < size; i++ { // func EncryptCBC(key, plaintext []byte) []byte {
		//	ciphertext := make([]byte, len(plaintext))
		//	return ciphertext
		// }

		result[i] = byte(r.Intn(256))
	}

	return result
}

func BlackBox(plaintext []byte) []byte {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	// FIXME: add padding before and after plaintext,
	// gives different errors for EBC anc CBC
	//bufSize := r.Intn(11)
	bufSize := 8
	if bufSize < 5 {
		bufSize += 5
	}
	buf := make([]byte, bufSize)
	plaintext = append(plaintext, buf...)
	plaintext = append(buf, plaintext...)
	key := RandomKey(16)
	iv := RandomKey(16)

	which := r.Intn(2)
	if which == 1 {
		fmt.Println("CBC encrypting ssh.")
		return EncryptCBC(key, plaintext, iv)
	}
	fmt.Println("ECB encrypting")
	return EncryptECB(key, plaintext)
}

func DetectECB(text []byte) bool {
	amt := len(text) / 16
	repeats := make(map[string]bool)
	for i := 0; i < amt; i++ {
		if repeats[string(text[i*16:i*16+16])] {
			return true
		}
		repeats[string(text[i*16:i*16+16])] = true
	}
	return false
}

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
	fmt.Println(string(PadPKCS([]byte("YELLOW SUBMARINE"), 20)))
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
	fmt.Println(ciphertext[:10])
	// initialization vector is all ASCII zeroes
	iv := make([]byte, len(key))
	plaintext := DecryptCBC(key, ciphertext, iv)
	fmt.Println(string(plaintext[:10])) // Play that funky music white boy
	ciphertext = EncryptCBC(key, plaintext, iv)
	fmt.Println(ciphertext[:10])
	plaintext = DecryptCBC(key, ciphertext, iv)
	fmt.Println(string(plaintext[:10]))
	// Challenge 11
	fmt.Println(DetectECB(BlackBox(plaintext)))

}
