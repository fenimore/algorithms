package main

import "fmt"

type HashValue struct {
	key   string
	value string
}

type HashTable struct {
	values [4][]string // max four items
}

func Put(key string, val string) {

}

func Hash(key string) int {
	if len(key) < 0 {
		return -1 // error
	}
	//fmt.Println(key[:1])
	//fmt.Println([]byte(key[:1])[0])
	//fmt.Println([]byte(key[:1])[0] >> 1)
	//fmt.Println("By two", []byte(key[:1])[0]<<6)
	//fmt.Printf("%b ", []byte(key[:1])[0]>>6)

	return int([]byte(key[:1])[0]) % 10
}

func main() {
	fmt.Println(" hello", Hash("hello"))
	fmt.Println(" hello", Hash("Hello"))
	fmt.Println("Brazen", Hash("brazen"))
	fmt.Println("zen", Hash("zen"))
	fmt.Println("Arazan", Hash("Arazen"))
	fmt.Println("2", Hash("2"))
	fmt.Println("3", Hash("3"))
	fmt.Println("4", Hash("4"))
	fmt.Println("5", Hash("5"))
	fmt.Println("1", Hash("1"))
	fmt.Println("0", Hash("0"))
	fmt.Println("-", Hash("-2"))
	fmt.Println("=", Hash("=2"))
}
