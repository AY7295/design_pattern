package main

import (
	"abstract_factory/furniture"
	"fmt"
)

func main() {

	f := furniture.NewFactory(furniture.ModernStyle, furniture.Wood)

	fmt.Printf("%+v", f.CreateChair().Information())

	fmt.Printf("%+v", f.CreateTable().Information())

}
