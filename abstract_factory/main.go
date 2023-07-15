package main

import (
	"abstract_factory/furniture"
	"fmt"
)

func main() {

	f := furniture.NewFactory(furniture.ModernStyle, furniture.Wood)

	c1 := f.CreateChair()
	fmt.Println(c1.Information())

	t1 := f.CreateTable()
	fmt.Println(t1.Information())

}
