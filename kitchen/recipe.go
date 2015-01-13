package kitchen

import (
	"github.com/casimir/storekeeper/storage"
)

type Recipe struct {
	ID          string
	Ingredients []storage.Stack
	Name        string
	Out         storage.Stack
}
