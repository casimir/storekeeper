package d3

import (
	"encoding/json"
	"log"

	"github.com/casimir/storekeeper/store"
	"github.com/casimir/storekeeper/util"
)

const (
	apiHostEU = "http://eu.battle.net"
	apiHostUS = "http://us.battle.net"

	apiArtisan = "/api/d3/data/artisan/"
	apiItem    = "/api/d3/data/item/"

	argLocale = "locale"

	StoreName = "D3"
)

var (
	_locales = []string{"en", "fr"}
)

type Provider struct {
	d3Artisans []Artisan
	itemQueue  util.StringSet
	store      *store.Store
}

func (p *Provider) Store() *store.Store {
	p.store = &store.Store{}
	p.initArtisans()
	p.initItems()
	return p.store
}

func (p *Provider) initArtisans() {
	p.store.Artisans = []store.Artisan{
		{"blacksmith", "Blacksmith", StoreName},
		{"jeweler", "Jeweler", StoreName},
	}

	f := store.Fetcher{apiHostEU + apiArtisan}
	for _, it := range p.store.Artisans {
		resp := f.Request(it.ID)
		if resp.Err != nil {
			log.Printf("Failed to get artisan information: %s", resp.Err)
			continue
		}
		var a Artisan
		err := json.Unmarshal(resp.Body, &a)
		if err != nil {
			log.Printf("Failed to get artisan information: %s", err)
			continue
		}
		p.store.Book = append(p.store.Book, a.ToBook(&p.itemQueue)...)
	}
}

func (p *Provider) initItems() {
	f := store.Fetcher{apiHostEU + apiItem}
	for _, resp := range f.Fetch(p.itemQueue.StringSlice) {
		if resp.Err != nil {
			log.Printf("Failed to get item information: %s", resp.Err)
			continue
		}
		var item Item
		err := json.Unmarshal(resp.Body, &item)
		if err != nil {
			log.Printf("Failed to get item information: %s", err)
			continue
		}
		p.store.Catalog = append(p.store.Catalog, item.normalize())
	}
}
