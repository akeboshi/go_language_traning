package main

func main() {
	println(impanic())
}

func impanic() (a int) {
	defer func() {
		recover()
		a = 1
	}()
	panic(0)
}
