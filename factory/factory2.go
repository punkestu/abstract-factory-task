package factory

import (
	"github.com/punkestu/abstract-factory-task/interfaces"
	"github.com/punkestu/abstract-factory-task/product"
)

type Factory2 struct {
}

func (f Factory2) Create(i int) interfaces.Product {
	if i == 2 {
		return product.Product2{}
	}
	if i == 4 {
		return product.Product4{}
	}
	return nil
}
