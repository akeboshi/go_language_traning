// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package bank

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)
var messages = make(chan string)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	Deposit(-amount)
	if Balance() < 0 {
		Deposit(amount)
		messages <- fmt.Sprintf("Failuer: your deposit is %d", Balance())
		<-messages
		return false
	}
	messages <- fmt.Sprintf("Success: your deposit is %d", Balance())
	<-messages
	return true
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case msg := <-messages:
			fmt.Println(msg)
			messages <- ""
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
