package d3

import (
	"encoding/json"
	"log"

	"github.com/casimir/storekeeper/store"
)

const (
	apiArtisan = "/api/d3/data/artisan/"
	apiData    = "/api/d3/data/"
	apiHostEU  = "http://eu.battle.net"
	apiHostUS  = "http://us.battle.net"

	argLocale = "locale"

	StoreName = "D3"
)

var (
	_locales = []string{"en", "fr"}
)

type D3Provider struct {
	d3Artisans []Artisan
	store      *store.Store
}

func (p *D3Provider) Store() *store.Store {
	p.store = &store.Store{}
	p.update()
	return p.store
}

func (p *D3Provider) update() error {
	p.initArtisans()

	for _, a := range p.d3Artisans {
		p.store.Book = append(p.store.Book, a.ToBook()...)
	}

	return nil
}

func (p *D3Provider) initArtisans() {
	p.store.Artisans = []store.Artisan{
		{"blacksmith", "Blacksmith", StoreName},
		{"jeweler", "Jeweler", StoreName},
	}

	f := store.Fetcher{}
	for _, a := range p.store.Artisans {
		resp := f.Request(apiHostEU + apiArtisan + a.Id)
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
		p.d3Artisans = append(p.d3Artisans, tmp)
	}
	return
}
