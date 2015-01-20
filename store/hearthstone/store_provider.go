package hearthstone

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/casimir/doable"
	"github.com/casimir/storekeeper/kitchen"
	"github.com/casimir/storekeeper/store"
)

const (
	apiURL = "http://hearthstonejson.com/json/AllSets.%s.json"

	DefaultLocale = "enUS"
	StoreName     = "Hearthstone"
)

var Locales = []string{"deDE", "enGB", "enUS", "esES", "esMX", "frFR", "itIT", "koKR", "plPL", "ptBR", "ptPT", "ruRU", "zhCN", "zhTW"}

// Card is the representation of an Hearthstone card.
type Card struct {
	// Name of the card.
	Name string
	// Cost is the mana cost of this card.
	Cost int
	// Type of the card. One of: Minion, Spell, Weapon, Hero, Hero Power,
	// Enchantment.
	Type string
	// The rarity of the card. Can be: Free, Common, Rare, Epic, Legendary
	// Note: Hearthstone internally uses 'Common' rarity on several cards in
	// the Basic set that are obtained for free. Thus these cards show a
	// 'Common' rarity even though the player gets them freely.
	Rarity string
	// Faction of the card. Can be: Alliance, Horde, Neutral.
	Faction string
	// Race of the card. Can be: Murloc, Demon, Beast, Totem, Pirate,
	// Dragon.
	Race string
	// The player class this card belongs to. Can be: Druid, Hunter, Mage,
	// Paladin.
	PlayerClass string
	// Text of the card when it is in your hand, in HTML.
	Text string
	// Text of the card when it is in play, in HTML.
	InPlayText string
	// The mechanics of the card. A combination of: Windfury, Combo, Secret,
	// Battlecry, Deathrattle, Taunt.
	Mechanics []string
	// The flavor text of the card.
	Flavor string
	// The artist of the card.
	Artist string
	// The attack of the card. Used for both Minions and Weapons.
	Attack int
	// The health of the card. Used for Minions.
	Health int
	// The durability of the card. Used for Weapons.
	Durability int
	// Hearthstone ID of the card.
	ID string
	// Whether this card can be acquired by the player or not.
	Collectible bool
	// Whether this card is elite or not.
	Elite bool
	// How to get this card. Only present if it's gotten via a method other
	// than opening a booster pack.
	HowToGet string
	// How to get the gold version of this card. Only present if it's gotten
	// via a method other than opening a booster pack.
	HowToGetGold string
}

func (c Card) UID() string {
	return c.ID
}

func (c Card) Match(other doable.Item) bool {
	o, ok := other.(Card)
	return ok && c.ID == o.ID
}

func (c Card) toRecipe() *kitchen.Recipe {
	r := &kitchen.Recipe{
		ID:   c.ID,
		Name: c.Name,
		Node: &doable.Node{
			Item: c,
			Nb:   1,
		},
	}
	r.Node.AddDep(&doable.Node{
		Item: Card{ID: "dust", Name: "Dust"},
		Nb:   c.Cost,
	})
	return r
}

type Provider struct {
	store *store.Store
}

func (p Provider) Store() *store.Store {
	p.store = store.NewStore()

	f := store.Fetcher{}
	r := f.Request(fmt.Sprintf(apiURL, DefaultLocale))
	if r.Err != nil {
		log.Fatalf("Failed to get cards: %s", r.Err)
	}
	var sets map[string][]Card
	err := json.Unmarshal(r.Body, &sets)
	if err != nil {
		log.Fatalf("Failed to decode cards: %s", err)
	}

	for set, cards := range sets {
		p.store.Artisans = append(p.store.Artisans, &store.Artisan{ID: set, Label: set})
		for _, it := range cards {
			p.store.Book = append(p.store.Book, it.toRecipe())
			p.store.Catalog.Add(it)
		}
	}

	return p.store
}
