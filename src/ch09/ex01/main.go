// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import "ch09/bank"

func main() {
	bank.Deposit(100)
	bank.Withdraw(50)
	bank.Withdraw(100)
}
