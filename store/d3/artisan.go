package d3

import (
	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/storage"
)

type (
	Artisan struct {
		Slug     string
		Name     string
		Portrait string
		Training map[string][]Tier
	}

	Tier struct {
		Tier   int
		Levels []Level
	}

	Level struct {
		Tier           int
		TierLevel      int
		Percent        int
		TrainedRecipes []Recipe
		TaughtRecipes  []Recipe
		UpgradeCost    int
	}

	Recipe struct {
		Id           string
		Slug         string
		Name         string
		Cost         int
		Reagents     []Stack
		ItemProduced Item
	}

	Stack struct {
		Quantity int
		Item     Item
	}

	Item struct {
		Id            string
		Name          string
		Icon          string
		DisplayColor  string
		TooltipParams string
	}
)

func (a Artisan) ToBook(items *StringSet) (book []kitchen.Recipe) {
	for _, tier := range a.Training["tiers"] {
		for _, lvl := range tier.Levels {
			for _, r := range lvl.TrainedRecipes {
				book = append(book, r.normalize())
				items.Add(r.ItemProduced.Id)
				for _, it := range r.Reagents {
					items.Add(it.Item.Id)
				}
			}
			for _, r := range lvl.TaughtRecipes {
				book = append(book, r.normalize())
				items.Add(r.ItemProduced.Id)
				for _, it := range r.Reagents {
					items.Add(it.Item.Id)
				}
			}
		}
	}
	return
}

func (i Item) normalize() storage.Item {
	return storage.Item{i.Id, i.Name}
}

func (r Recipe) normalize() kitchen.Recipe {
	ret := kitchen.Recipe{
		ID:          r.Id,
		Ingredients: []storage.Stack{},
		Name:        r.Name,
		Out:         storage.Stack{1, r.ItemProduced.normalize()},
	}

	for _, it := range r.Reagents {
		s := storage.Stack{it.Quantity, it.Item.normalize()}
		ret.Ingredients = append(ret.Ingredients, s)
	}

	return ret
}
