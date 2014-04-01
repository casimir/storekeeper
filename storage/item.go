package storage

type Item struct {
	Id   string
	Name string
}

type Stack struct {
	Count int
	Item  Item
}
