package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/casimir/storekeeper/util"
)

var dbDir = util.ApplicationPath + "/data"

func dbPath(name string) string {
	return fmt.Sprintf("%s/%s.json", dbDir, name)
}

// Reserve provides a way to persist a Store.
type Reserve struct {
	name string
}

// NewReserve create a reserve with the given name.
func NewReserve(name string) *Reserve {
	os.MkdirAll(dbDir, 0755)
	return &Reserve{name: name}
}

// Load the Store from the Reserve.
func (r *Reserve) Load() *Store {
	ret := new(Store)
	data, err := ioutil.ReadFile(dbPath(r.name))
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
	if err := ioutil.WriteFile(dbPath(r.name), data, 0644); err != nil {
		return fmt.Errorf("Failed to set %s: %s", r.name, err)
	}
	return nil
}

// DeleteReserve delete a persisted store. Does nothing if the store does not
// exist.
func DeleteReserve(name string) { os.Remove(dbPath(name)) }
