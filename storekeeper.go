package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/casimir/storekeeper/store/d3"
)

const (
	StoreD3 = iota
)

func getStore(storeID int) (Store, error) {
	switch storeID {
	case StoreD3:
		return new(d3.Store), nil
	}
	return nil, errors.New("Unknown store type")
}

func main() {
	s, _ := getStore(StoreD3)
	if err := s.Update(nil); err != nil {
		fmt.Printf("Init failed: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Book loaded - %d items\n", len(s.Book()))
}
