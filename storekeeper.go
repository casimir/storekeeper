package main

import (
    "fmt"
    "github.com/chibibi/storekeeper/store/d3"
    "os"
)

const (
    StoreD3 = iota
)

func getStore(storeID int) Store {
    switch storeID {
    case StoreD3:
        return new(d3.Store)
    }
    return nil
}

func main() {
    s := getStore(StoreD3)
    if err := s.Init(nil); err != nil {
        fmt.Printf("Init failed: %s\n", err.Error())
        os.Exit(1)
    }
    fmt.Printf("Book loaded - %d items\n", len(s.Book()))
}
