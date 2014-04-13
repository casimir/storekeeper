package store

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/user"

	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/storage"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Reserve struct {
	dbm *gorp.DbMap
}

func ApplicationPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir + "/.storekeeper"
}

func (r *Reserve) Init(storeName string) {
	dbPath := ApplicationPath() + "/db/" + storeName + ".db"
	os.Remove(dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	r.dbm = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}, TypeConverter: converter{}}

	r.dbm.AddTableWithName(Artisan{}, "artisans").SetKeys(false, "Id")
	r.dbm.AddTableWithName(storage.Item{}, "items").SetKeys(false, "Id")
	r.dbm.AddTableWithName(kitchen.Recipe{}, "recipes").SetKeys(false, "Id")

	if err := r.dbm.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}
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
			log.Printf("Error while saving artisans: %s", err.Error())
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
			log.Printf("-> %s", s)
			return json.Unmarshal([]byte(*s), target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	}
	return gorp.CustomScanner{}, false
}
