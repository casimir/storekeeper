package kitchen

import (
	"errors"

	"github.com/casimir/doable"
)

type (
	Recipe struct {
		ID   string
		Name string
		Node *doable.Node
	}

	Cook struct {
		recipe *Recipe
		stock  *doable.List
	}
)

func NewCook(recipe *Recipe, stock *doable.List) *Cook {
	return &Cook{
		recipe: recipe,
		stock:  stock,
	}
}

func (c Cook) Cook() error {
	t := doable.New(c.recipe.Node, c.stock)
	if !t.Doable() {
		return errors.New("cook: not enough ingredients for this recipe")
	}
	return nil
}

func (c Cook) IsCookable() bool {
	t := doable.New(c.recipe.Node, c.stock.Clone())
	return t.Doable()
}

func (c Cook) MissingIngredients() *doable.List {
	t := doable.New(c.recipe.Node, c.stock.Clone())
	if t.Doable() {
		return doable.NewList()
	}
	return t.Miss
}
