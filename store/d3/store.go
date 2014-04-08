package d3

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/storage"
	"github.com/casimir/storekeeper/store"
)

const (
	apiArtisan = "/api/d3/data/artisan/"
	apiData    = "/api/d3/data/"
	apiHostEU  = "http://eu.battle.net"
	apiHostUS  = "http://us.battle.net"

	argLocale = "locale"
)

var (
	_artisans = []string{"blacksmith", "jeweler"}
	_locales  = []string{"en", "fr"}
)

type Store struct {
	artisans []Artisan
	book     []kitchen.Recipe
	catalog  []storage.Item
}

func (s Store) Bag() (*storage.Bag, error) {
	return nil, errors.New("Not implemented yet")
}

func (s Store) Book() []kitchen.Recipe {
	return s.book
}

func (s Store) Catalog() []storage.Item {
	return s.catalog
}

func (s *Store) Update(args map[string]string) error {
	s.initArtisans()
	s.initItems()

	for _, a := range s.artisans {
		s.book = append(s.book, a.ToBook()...)
		s.catalog = append(s.catalog, a.ToCatalog()...)
	}

	return nil
}

func (s *Store) initArtisans() {
	f := store.Fetcher{}
	for _, a := range _artisans {
		resp := f.Request(apiHostEU + apiArtisan + a)
		if resp.Err != nil {
			log.Printf("Failed to get artisan: %s", resp.Err)
			continue
		}
		var tmp Artisan
		err := json.Unmarshal(resp.Body, &tmp)
		if err != nil {
			log.Printf("Failed to get artisan: %s", err)
			continue
		}
		s.artisans = append(s.artisans, tmp)
		log.Printf("Got artisan: %s", a)
	}
	return
}

func (s *Store) initItems() {
	// TODO check uniq
}
