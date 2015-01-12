package store

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/storage"
	"github.com/casimir/storekeeper/util"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3" // SQlite connection driver
)

var dbDir = util.ApplicationPath() + "/db"

type Reserve struct {
	dbm *gorp.DbMap
}

func init() {
	os.MkdirAll(dbDir, 0755)
}

func dbPath(name string) string {
	return dbDir + "/" + name + ".db"
}

func DeleteReserve(storeName string) {
	os.Remove(dbPath(storeName))
}

func NewReserve(storeName string) *Reserve {
	db, err := sql.Open("sqlite3", dbPath(storeName))
	if err != nil {
		panic(err)
	}
	ret := &Reserve{
		&gorp.DbMap{
			Db:            db,
			Dialect:       gorp.SqliteDialect{},
			TypeConverter: converter{},
		},
	}
	if ret.init() != nil {
		panic(err)
	}
	return ret
}

func (r *Reserve) init() error {
	r.dbm.AddTableWithName(Artisan{}, "artisans").SetKeys(false, "Id")
	r.dbm.AddTableWithName(storage.Item{}, "items").SetKeys(false, "Id")
	r.dbm.AddTableWithName(kitchen.Recipe{}, "recipes").SetKeys(false, "Id")
	return r.dbm.CreateTablesIfNotExists()
}

func (r Reserve) Load() *Store {
	s := &Store{}
	_, err := r.dbm.Select(&s.Artisans, "select * from artisans")
	if err != nil {
		log.Printf("Failed to load artisans: %s\n", err)
		err = nil
	}
	_, err = r.dbm.Select(&s.Book, "select * from recipes")
	if err != nil {
		log.Printf("Failed to load recipes: %s\n", err)
		err = nil
	}
	_, err = r.dbm.Select(&s.Catalog, "select * from items")
	if err != nil {
		log.Printf("Failed to load items: %s\n", err)
		err = nil
	}
	return s
}

func (r Reserve) Save(s *Store) {
	for _, a := range s.Artisans {
		if err := r.dbm.Insert(&a); err != nil {
			log.Printf("Error while saving artisans: %s", err.Error())
		}
	}
	for _, b := range s.Book {
		if err := r.dbm.Insert(&b); err != nil {
			log.Printf("Error while saving recipes: %s", err.Error())
		}
	}
	for _, c := range s.Catalog {
		if err := r.dbm.Insert(&c); err != nil {
			log.Printf("Error while saving catalog: %s", err.Error())
		}
	}
}

type converter struct{}

func (c converter) ToDb(val interface{}) (interface{}, error) {
	switch v := val.(type) {
	case storage.Stack, []storage.Stack:
		raw, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(raw), nil
	}
	return val, nil
}

func (c converter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *storage.Stack, *[]storage.Stack:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("Error while loading Stack type")
			}
			return json.Unmarshal([]byte(*s), target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	}
	return gorp.CustomScanner{}, false
}
