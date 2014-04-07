package kitchen

import (
	"errors"

	"github.com/casimir/storekeeper/storage"
)

type Cook struct {
	bag *storage.Bag
}

func (c Cook) Cook(r Recipe) error {
	if !c.IsCookable(r) {
		return errors.New("cook: not enough ingredients for this recipe")
	}

	for _, v := range r.Ingredients {
		c.bag.RemoveItem(v.Count, v.Item.Id)
	}
	c.bag.AddItem(r.Out.Count, r.Out.Item.Id)
	return nil
}

func (c Cook) IsCookable(r Recipe) bool {
	for _, v := range r.Ingredients {
		if v.Count > c.bag.Count(v.Item.Id) {
			return false
		}
	}
	return true
}

func (c Cook) MissingIngredients(r Recipe) []storage.Stack {
	if c.IsCookable(r) {
		return []storage.Stack{}
	}

	ret := []storage.Stack{}
	for _, v := range r.Ingredients {
		if missing := v.Count - c.bag.Count(v.Item.Id); missing > 0 {
			v.Count = missing
			ret = append(ret, v)
		}
	}
	return ret
}
