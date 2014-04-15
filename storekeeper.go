package main

import (
	"log"

	"github.com/casimir/storekeeper/store"
	"github.com/casimir/storekeeper/store/d3"
)

func main() {
	p := d3.D3Provider{}
	s := p.Store()
	r := &store.Reserve{}
	r.Init(d3.StoreName)
	log.Print("Saving data...")
	r.Save(s)
	log.Print("Store saved to database")

	log.Print("loading data...")
	loadedStore := r.Load()
	log.Println("Store loaded from database")
	log.Printf("- %d artisans\n", len(loadedStore.Artisans))
	log.Printf("- %d recipes\n", len(loadedStore.Book))
	log.Printf("- %d items\n", len(loadedStore.Catalog))
}
