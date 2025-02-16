package cache

import "github.com/k-ksu/avito-shop/internal/model"

type Merch map[string]model.Merch

func NewMerch() Merch {
	return make(map[string]model.Merch)
}

func (merch Merch) Add(item model.Merch) {
	merch[item.Name] = item
}

func (merch Merch) Get(name string) (model.Merch, bool) {
	m, ok := merch[name]
	return m, ok
}
