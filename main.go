package main

import "design-patterns/factory"

func main() {
	//instance := singleton.GetInstance()
	//fmt.Println(instance.Data)

	factoryA := factory.ConCreateFactoryA{}
	productA := factoryA.Create()
	productA.ShowInfo()

	factoryB := factory.ConCreateFactoryB{}
	productB := factoryB.Create()
	productB.ShowInfo()
}
