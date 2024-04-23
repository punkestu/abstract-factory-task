package interfaces

type Product interface {
	CallMe() string
}

type Factory interface {
	Create(int) Product
}
