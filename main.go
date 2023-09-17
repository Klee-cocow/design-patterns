package main

import "design-patterns/adapter"

func main() {
	//单例
	//instance := singleton.GetInstance()
	//fmt.Println(instance.Data)

	//工厂
	//factoryA := factory.ConCreateFactoryA{}
	//productA := factoryA.Create()
	//productA.ShowInfo()
	//
	//factoryB := factory.ConCreateFactoryB{}
	//productB := factoryB.Create()
	//productB.ShowInfo()

	//建造者模式
	//build := &builder.ConCreateBuild{}
	//director := builder.NewDirector(build)
	//director.Construct()
	//result := build.GetResult()
	//
	//fmt.Printf("构建的产品包括: ,%s ,%s ,%s", result.Party1, result.Party2, result.Party3)

	ukPlug := adapter.UKPlug{}
	plugAdapter := adapter.UKPlugAdapter{Plug: ukPlug}
	plugAdapter.InsertChineseSocket()
}
