package main

import (
	"fmt"

	"github.com/punkestu/abstract-factory-task/factory"
	"github.com/punkestu/abstract-factory-task/interfaces"
)

type AbstractFactory struct {
}

func (af AbstractFactory) Create(i string) interfaces.Factory {
	if i == "factory1" {
		return factory.Factory1{}
	}
	if i == "factory2" {
		return factory.Factory2{}
	}
	return nil
}

func main() {
	var af AbstractFactory

	fmt.Println("+Creating factory1")
	var f1 interfaces.Factory = af.Create("factory1")
	if f1 == nil {
		fmt.Println("-Factory not found")
	} else {
		fmt.Println("++Creating product1")
		var p1 interfaces.Product = f1.Create(1)
		if p1 == nil {
			fmt.Println("--Product not found")
		} else {
			fmt.Println("+++Call product1")
			fmt.Println(">>>" + p1.CallMe())
		}
		fmt.Println("++Creating product3")
		var p3 interfaces.Product = f1.Create(3)
		if p3 == nil {
			fmt.Println("--Product not found")
		} else {
			fmt.Println("+++Call product3")
			fmt.Println(">>>" + p3.CallMe())
		}
	}

	fmt.Println()

	fmt.Println("+Creating factory2")
	var f2 interfaces.Factory = af.Create("factory2")
	if f2 == nil {
		fmt.Println("-Factory not found")
	} else {
		fmt.Println("++Creating product2")
		var p2 interfaces.Product = f2.Create(2)
		if p2 == nil {
			fmt.Println("--Product not found")
		} else {
			fmt.Println("+++Call product2")
			fmt.Println(">>>" + p2.CallMe())
		}
		fmt.Println("++Creating product4")
		var p4 interfaces.Product = f2.Create(4)
		if p4 == nil {
			fmt.Println("--Product not found")
		} else {
			fmt.Println("+++Call product4")
			fmt.Println(">>>" + p4.CallMe())
		}
	}
}
