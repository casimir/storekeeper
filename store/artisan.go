package store

type Artisan struct {
	ID     string `db:"Id"`
	Label  string
	Source string
}
