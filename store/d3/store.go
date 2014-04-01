package d3

import (
    "encoding/json"
    "github.com/chibibi/storekeeper/kitchen"
    "github.com/chibibi/storekeeper/storage"
    "io/ioutil"
    "log"
    "net/http"
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

func (s Store) Bag() *storage.Bag {
    return nil
}

func (s Store) Book() []kitchen.Recipe {
    return s.book
}

func (s Store) Catalog() []storage.Item {
    return s.catalog
}

func (s *Store) Init(args map[string]string) error {
    s.initArtisans()
    s.initItems()

    for _, a := range s.artisans {
        s.book = append(s.book, a.ToBook()...)
        s.catalog = append(s.catalog, a.ToCatalog()...)
    }

    return nil
}

func (s *Store) initArtisans() {
    for _, a := range _artisans {
        resp, err := http.Get(apiHostEU + apiArtisan + a)
        if err != nil {
            log.Printf("Failed to get artisan: %s", err)
            continue
        }
        defer resp.Body.Close()
        raw, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Printf("Failed to get artisan: %s", err)
            continue
        }
        var tmp Artisan
        err = json.Unmarshal(raw, &tmp)
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
