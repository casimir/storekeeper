package main

import (
	"log"

	"github.com/casimir/storekeeper/store"
	"github.com/casimir/storekeeper/store/d3"
)

func main() {
	p := d3.D3Provider{}
	log.Print("Fetching data...")
	s := p.Store()
	log.Print("Store fetched")
	r := store.NewReserve(d3.StoreName)
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
