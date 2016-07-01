//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

func main() {
	println(impanic())
}

func impanic() (a int) {
	a = 1
	defer func() {
		recover()
	}()
	panic(0)
}
