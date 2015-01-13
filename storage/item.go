package storage

type Item struct {
	ID   string
	Name string
}

type Stack struct {
	Count int
	Item  Item
}
