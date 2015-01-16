package storage

import (
	"errors"
)

type (
	Item struct {
		ID   string
		Name string
	}

	Stack struct {
		Count int
		Item  Item
	}

	Bag struct {
		Items map[string]*Stack
	}
)

func (b *Bag) AddItem(count int, id string) error {
	if b.Items == nil {
		b.Items = make(map[string]*Stack)
	}

	if item, ok := b.Items[id]; !ok {
		b.Items[id] = &Stack{Count: count}
	} else {
		item.Count += count
	}
	return nil
}

func (b Bag) Count(id string) (n int) {
	for k, v := range b.Items {
		if k == id {
			n += v.Count
		}
	}
	return n
}

func (b *Bag) RemoveItem(count int, id string) error {
	item, ok := b.Items[id]
	if !ok {
		return errors.New("No item in the bag")
	}
	if item.Count < count {
		return errors.New("Not enough item in the bag")
	}

	item.Count -= count
	if item.Count == 0 {
		delete(b.Items, id)
	}
	return nil
}
