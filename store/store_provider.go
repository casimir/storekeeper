package store

type Provider interface {
	Store() (*Store, error)
}
