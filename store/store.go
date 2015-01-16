package store

import (
	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/storage"
)

type Provider interface {
	Store() *Store
}

type Store struct {
	Artisans []Artisan
	Book     []kitchen.Recipe
	Catalog  []storage.Item
}
