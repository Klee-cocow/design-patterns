package adapter

import "fmt"

// UKPlug 表示英国插头
type UKPlug struct{}

func (p UKPlug) InsertUKSocket() {
	fmt.Println("插入英国插座")
}

// ChineseSocket 表示中国插座接口
type ChineseSocket interface {
	InsertChineseSocket()
}

// UKPlugAdapter 是适配器，将英国插头适配到中国插座
type UKPlugAdapter struct {
	Plug UKPlug
}

func (a UKPlugAdapter) InsertChineseSocket() {
	fmt.Println("使用适配器将英国插头适配到中国插座")
	a.Plug.InsertUKSocket()
}
