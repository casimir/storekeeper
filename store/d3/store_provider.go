package d3

import (
	"encoding/json"
	"log"

	"github.com/casimir/storekeeper/store"
)

const (
	apiHostEU = "http://eu.battle.net"
	apiHostUS = "http://us.battle.net"

	apiArtisan = "/api/d3/data/artisan/"
	apiItem    = "/api/d3/data/item/"

	StoreName = "D3"
)

type Provider struct {
	d3Artisans []Artisan
	itemQueue  StringSet
	store      *store.Store
}

func (p Provider) Store() *store.Store {
	p.store = new(store.Store)
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
		r := f.Request(it.ID)
		if r.Err != nil {
			log.Printf("Failed to get artisan information: %s", r.Err)
			continue
		}
		var a Artisan
		err := json.Unmarshal(r.Body, &a)
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
