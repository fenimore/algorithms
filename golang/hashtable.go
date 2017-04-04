// Hash table implementation
// for simplicity, only use int as key
// TODO: convert string to int for string as key?
// get/set/remove
package main

import "fmt"
import "errors"

type value struct {
	key string
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
func (t *table) hash(key string) int {
	// mod keeps the value in the correct range
	val := key[0] + key[1]
	return int(val) % t.size
}

func (t *table) set(key string, val string) {
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

func (t *table) get(key string) (string, error) {
	for _, x := range t.table[t.hash(key)] {
		if x.key == key {
			return x.val, nil
		}
	}
	// if doesn't exist, return error
	return "", errors.New("Error Get: Key doesn't exist")
}

func (t *table) rem(key string) error {
	hash := t.hash(key)
	for _, x := range t.table[hash] {
		if x.key == key {
			// delete item from list
			// TODO: allow for collisions
			t.table[hash] = make([]value, 0)
			return nil
		}
	}
	return errors.New("Error Rem: Key doesn't exist")
}

func main() {
	var err error
	t := NewTable(128)
	t.set("foo", "testing")
	fmt.Println(t.get("woot"))
	t.set("bar", "Woot")
	fmt.Println(t.get("foo"))
	//  deletion
	fmt.Println("\nDeletion:")
	fmt.Println(t.get("bar"))
	err = t.rem("bar")
	fmt.Println(err)
	fmt.Println(t.get("bar"))
	err = t.rem("bar")
	fmt.Println(err)
	// TODO: check for collisions

}
