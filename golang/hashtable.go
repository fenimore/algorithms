// Hash table implementation
// for simplicity, only use int as key
// TODO: convert string to int for string as key?
// get/set/remove
package main

import "fmt"
import "errors"

type value struct {
	key int
	val string
}

type table struct {
	size  int
	table [][]value
}

func NewTable(size int) *table {
	t := new(table)
	t.size = size
	t.table = make([][]value, size)
	return t
}

// hash only takes keys that are ints
func (t *table) hash(key int) int {
	// mod keeps the value in the correct range
	return key % t.size
}

func (t *table) set(key int, val string) {
	idx := t.hash(key)
	for _, x := range t.table[idx] {
		if x.key == key {
			x.val = val
			return
		}
	}
	// if doesn't exist, add to table at key
	t.table[idx] = append(t.table[idx], value{key, val})
}

func (t *table) get(key int) (string, error) {
	for _, x := range t.table[t.hash(key)] {
		if x.key == key {
			return x.val, nil
		}
	}
	// if doesn't exist, return error
	return "", errors.New("Key doesn't exist")
}

func (t *table) rem(key int) error {
	hash := t.hash(key)
	for _, x := range t.table[hash] {
		if x.key == key {
			// delete item from list
			if len(t.table[hash]) > 1 {
				t.table[hash] = append(t.table[hash][:key], t.table[hash][key+1:]...)
			} else {
				t.table[hash] = make([]value, 0)
			}
			return nil
		}
	}
	return errors.New("Key doesn't exist")
}

func main() {
	var err error
	t := NewTable(128)
	t.set(64, "testing")
	fmt.Println(t.get(64))
	t.set(1, "Woot")
	fmt.Println(t.get(2))
	//  deletion
	fmt.Println("\nDeletion:")
	fmt.Println(t.get(1))
	err = t.rem(1)
	fmt.Println(err)
	fmt.Println(t.get(1))
}
