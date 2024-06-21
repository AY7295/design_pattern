package main

import "singleton/register"

func main() {
	register.Global.Store("key", "value")
	value, ok := register.Global.Load("key")
	if ok {
		println(value.(string))
	}

}
