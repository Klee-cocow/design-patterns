package factory

import "fmt"

type Product interface {
	ShowInfo()
}

type Factory interface {
	Create() Product
}

type ConCreateProductA struct {
}

type ConCreateProductB struct {
}

type ConCreateFactoryA struct {
}
type ConCreateFactoryB struct {
}

func (p ConCreateProductA) ShowInfo() {
	fmt.Println("这是产品A")
}

func (p ConCreateProductB) ShowInfo() {
	fmt.Println("这是产品B")
}

func (p ConCreateFactoryA) Create() Product {
	return ConCreateProductA{}
}

func (p ConCreateFactoryB) Create() Product {
	return ConCreateProductB{}
}
