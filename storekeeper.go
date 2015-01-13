package main

import (
	"log"

	"github.com/casimir/storekeeper/store"
	"github.com/casimir/storekeeper/store/d3"
	"github.com/casimir/storekeeper/store/hearthstone"
)

func test(p store.Provider, name string) {
	log.Println(name)
	log.Print("Fetching data...")
	s := p.Store()
	log.Print("Store fetched")
	r := store.NewReserve(name)
	log.Print("Saving data...")
	r.Save(s)
	log.Print("Store saved to database")

	log.Print("Loading data...")
	loadedStore := r.Load()
	log.Println("Store loaded from database")
	log.Printf("- %d artisans\n", len(loadedStore.Artisans))
	log.Printf("- %d recipes\n", len(loadedStore.Book))
	log.Printf("- %d items\n", len(loadedStore.Catalog))
}

func main() {
	test(d3.Provider{}, d3.StoreName)
	test(hearthstone.Provider{}, hearthstone.StoreName)
}
