package storage

type Item struct {
	ID   string `db:"Id"`
	Name string
}

type Stack struct {
	Count int
	Item  Item
}
