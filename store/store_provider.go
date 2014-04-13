package store

type StoreProvider interface {
	Store() (*Store, error)
}
