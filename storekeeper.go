package main

import (
	"log"
	"os"
	"runtime"

	"github.com/casimir/storekeeper/store"
	"github.com/casimir/storekeeper/store/d3"
	"github.com/casimir/storekeeper/store/hearthstone"
)

var ApplicationPath string

func init() {
	switch runtime.GOOS {
	case "darwin":
		ApplicationPath = os.Getenv("HOME") + "/Library/Application Support/Storekeeper/"
	case "windows":
		ApplicationPath = os.Getenv("APPDATA") + "/Storekeeper/"
	default:
		ApplicationPath = os.Getenv("HOME") + "/.storekeeper/"
	}
}

func test(p store.Provider, name string) {
	log.Println(name)
	log.Print("Fetching data...")
	s := p.Store()
	log.Print("Store fetched")
	r := store.NewReserve(ApplicationPath, name)
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
