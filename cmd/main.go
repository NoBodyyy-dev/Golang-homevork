package main

import (
	"Golang-homevork/internal/pkg/storage"
	"fmt"
)

func main() {
	s, err := storage.NewStructure()
	if err != nil {
		s.Logger.Fatal(err.Error())
	}
	s.Set("key1", 1)
	s.Set("key2", "1")
	s.Set("key3", true)
	fmt.Println(*s.Get("key1"))
	fmt.Println(*s.GetKind("key2"))
}
