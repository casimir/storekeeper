package main

import (
	"github.com/chibibi/storekeeper/kitchen"
	"github.com/chibibi/storekeeper/storage"
)

type Store interface {
	Bag() *storage.Bag
	Book() []kitchen.Recipe
	Catalog() []storage.Item
	Init(args map[string]string) error
}
