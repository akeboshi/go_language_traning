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
