package main

import (
	"fmt"
	"math/rand"

	"github.com/timtadh/data-structures/tree/bptree"
	"github.com/timtadh/data-structures/types"
)

func main() {

	recs := make(records, 100)
	ranrec := func() *record {
		return &record{RandStringRunes(20), RandomNumber(100)}
	}

	for i := range recs {
		recs[i] = ranrec()
	}

	t := bptree.NewBpTree(23)

	for _, r := range recs {
		t.Add(r.key, r.value)
	}
	for _, r := range recs {
		fmt.Println(r.key, r.value, t.Has(r.key))
	}

	fmt.Println(t)

}

func RandomNumber(n int) types.Int {
	return types.Int(rand.Intn(n))
}

type record struct {
	key   types.String
	value types.Int
}

type records []*record

func (r records) Len() int {
	return len(r)
}

func (r records) Less(i, j int) bool {
	return r[i].key.Less(r[j].key)
}

func (r records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// func randstr(length int) types.String {
// 	return types.String(randslice(length))
// }
//
// func randslice(length int) []byte {
// 	if urandom, err := os.Open("/dev/urandom"); err != nil {
// 		panic(err)
// 	} else {
// 		slice := make([]byte, length)
// 		if _, err := urandom.Read(slice); err != nil {
// 			panic(err)
// 		}
// 		urandom.Close()
// 		// return append([]byte("b"), slice...)
// 		return slice
// 	}
// 	panic("unreachable")
// }
//
// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) types.String {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return types.String(string(b))
}
