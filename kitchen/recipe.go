package kitchen

import (
	"github.com/casimir/storekeeper/storage"
)

type Recipe struct {
	ID          string `db:"Id"`
	Ingredients []storage.Stack
	Name        string
	Out         storage.Stack
}
