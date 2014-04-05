package main

import (
	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/storage"
)

type Store interface {
	Bag() *storage.Bag
	Book() []kitchen.Recipe
	Catalog() []storage.Item
	Init(args map[string]string) error
}
