package main

import (
	"example.com/m/taodb"
	"fmt"
)

func main() {
    db, err := taodb.Open(taodb.DefaultOptions)
	if err != nil{
		panic(err)
	}

	db.Put([]byte("hello"), []byte("world"))
	fmt.Println(db.Get([]byte("hello")))
	fmt.Println(db.Get([]byte("hello_")))
}