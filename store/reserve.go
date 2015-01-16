package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Reserve provides a way to persist a Store.
type Reserve struct {
	name string
	path string
}

// NewReserve create a reserve with the given name.
func NewReserve(appPath, name string) *Reserve {
	dbPath := appPath + "data/"
	os.MkdirAll(dbPath, 0755)
	return &Reserve{
		name: name,
		path: dbPath + name + ".json",
	}
}

// Delete the persisted store. Does nothing if the db does not exist.
func (r Reserve) Delete() {
	os.Remove(r.path)
}

// Load the Store from the Reserve.
func (r *Reserve) Load() *Store {
	ret := new(Store)
	data, err := ioutil.ReadFile(r.path)
	if err != nil {
		log.Printf("Failed to load store %s: %s\n", r.name, err)
		return &Store{}
	}
	json.Unmarshal(data, &ret)
	return ret
}

// Save the store to the Reserve.
func (r *Reserve) Save(s *Store) error {
	data, _ := json.Marshal(s)
	if err := ioutil.WriteFile(r.path, data, 0644); err != nil {
		return fmt.Errorf("Failed to set %s: %s", r.name, err)
	}
	return nil
}
