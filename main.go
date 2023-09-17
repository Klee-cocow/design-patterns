package main

import (
	"design-patterns/singleton"
	"fmt"
)

func main() {
	instance := singleton.GetInstance()
	fmt.Println(instance.Data)
}
