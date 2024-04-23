package factory

import (
	"github.com/punkestu/abstract-factory-task/interfaces"
	"github.com/punkestu/abstract-factory-task/product"
)

type Factory1 struct {
}

func (f Factory1) Create(i int) interfaces.Product {
	if i == 1 {
		return product.Product1{}
	}
	if i == 3 {
		return product.Product3{}
	}
	return nil
}
