package main

import "fmt"

// Product is the interface for the objects the factory will create
type Product interface {
	Use()
}

// ConcreteProductA is a type of product the factory can create
type ConcreteProductA struct{}

// Use implements the Product interface
func (p *ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProductA")
}

// ConcreteProductB is a type of product the factory can create
type ConcreteProductB struct{}

// Use implements the Product interface
func (p *ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProductB")
}

// Factory is the interface for the object factory
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA is a type of factory that creates ConcreteProductA objects
type ConcreteFactoryA struct{}

// CreateProduct implements the Factory interface
func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB is a type of factory that creates ConcreteProductB objects
type ConcreteFactoryB struct{}

// CreateProduct implements the Factory interface
func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// Create a ConcreteFactoryA and use it to create a ConcreteProductA
	factory := &ConcreteFactoryA{}
	product := factory.CreateProduct()
	product.Use()

	// Create a ConcreteFactoryB and use it to create a ConcreteProductB
	factory = &ConcreteFactoryB{}
	product = factory.CreateProduct()
	product.Use()
}
