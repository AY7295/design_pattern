package main

import (
	"builder/house"
	"fmt"
)

func main() {
	sv := house.NewSupervisor()
	sv.SetBuilders(house.CustomLockBuilder("copper", "yellow", 2), house.CustomWindowBuilder("glass", "transport", 6))

	fmt.Println(sv.BuildHouse().Information())
}
