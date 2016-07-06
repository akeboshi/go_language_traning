// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

func main() {
	in, out := pipeline(100)
	go func() { in <- struct{}{} }()
	<-out
}

func pipeline(size int) (chan<- struct{}, <-chan struct{}) {
	var out chan struct{}

	c0 := make(chan struct{})
	out = c0

	for i := 0; i < size; i++ {
		in := make(chan struct{})
		go routine(in, out)
		out = in
	}
	return c0, out
}

func routine(in chan<- struct{}, out <-chan struct{}) {
	for {
		in <- <-out
	}
}
