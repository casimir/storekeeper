package kitchen

import (
	"github.com/chibibi/storekeeper/storage"
)

type Recipe struct {
	Id          string
	Ingredients []storage.Stack
	Name        string
	Out         storage.Stack
}
