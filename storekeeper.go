package main

import (
	"log"
	"os"
	"os/user"

	"github.com/casimir/storekeeper/store"
	"github.com/casimir/storekeeper/store/d3"
)

func init() {
	os.Mkdir(ApplicationPath(), 0755)
	os.Mkdir(ApplicationPath()+"/db", 0755)
}

func ApplicationPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir + "/.storekeeper"
}

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
