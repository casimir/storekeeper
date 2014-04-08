package main

import (
	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/storage"
)

type Store interface {
	Bag() (*storage.Bag, error)
	Book() []kitchen.Recipe
	Catalog() []storage.Item
	Update(args map[string]string) error
}
