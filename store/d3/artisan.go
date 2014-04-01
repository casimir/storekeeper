package d3

import (
    "github.com/chibibi/storekeeper/kitchen"
    "github.com/chibibi/storekeeper/storage"
)

type Artisan struct {
    Slug     string
    Name     string
    Portrait string
    Training map[string][]Tier
}

type Tier struct {
    Tier   int
    Levels []Level
}

type Level struct {
    Tier           int
    TierLevel      int
    Percent        int
    TrainedRecipes []Recipe
    TaughtRecipes  []Recipe
    UpgradeCost    int
}

type Recipe struct {
    Id           string
    Slug         string
    Name         string
    Cost         int
    Reagents     []Stack
    ItemProduced Item
}

type Stack struct {
    Quantity int
    Item     Item
}

type Item struct {
    Id            string
    Name          string
    Icon          string
    DisplayColor  string
    TooltipParams string
}

func (a Artisan) ItemData() (data []string) {
    for _, tier := range a.Training["tiers"] {
        for _, lvl := range tier.Levels {
            for _, r := range lvl.TrainedRecipes {
                for _, s := range r.Reagents {
                    data = append(data, s.Item.TooltipParams)
                }
                data = append(data, r.ItemProduced.TooltipParams)
            }
            for _, r := range lvl.TaughtRecipes {
                for _, s := range r.Reagents {
                    data = append(data, s.Item.TooltipParams)
                }
                data = append(data, r.ItemProduced.TooltipParams)
            }
        }
    }
    return
}

func (a Artisan) ToBook() (book []kitchen.Recipe) {
    for _, tier := range a.Training["tiers"] {
        for _, lvl := range tier.Levels {
            for _, r := range lvl.TrainedRecipes {
                book = append(book, kitchen.Recipe{
                    Id:   r.Id,
                    Name: r.Name,
                })
            }
            for _, r := range lvl.TaughtRecipes {
                book = append(book, kitchen.Recipe{
                    Id:   r.Id,
                    Name: r.Name,
                })
            }
        }
    }
    return
}

func (a Artisan) ToCatalog() (catalog []storage.Item) {
    return
}
