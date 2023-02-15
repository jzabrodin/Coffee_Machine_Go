package main

import (
	"fmt"
)

type Available struct {
	waterAvailable  int
	milkAvailable   int
	coffeeAvailable int
	cupsAvailable   int
	moneyAvailable  int
}

type Receipt struct {
	waterAmmount  int
	milkAmmount   int
	coffeeAmmount int
	moneyAmmount  int
}

var available Available

func Espresso() Receipt {
	// For one espresso, the coffee machine needs 250 ml of water and 16 g of coffee beans. It costs $4.
	var espresso Receipt
	espresso.coffeeAmmount = 16
	espresso.milkAmmount = 0
	espresso.moneyAmmount = 4
	espresso.waterAmmount = 250

	return espresso
}

func Latte() Receipt {
	//For a latte, the coffee machine needs 350 ml of water, 75 ml of milk, and 20 g of coffee beans. It costs $7.
	var espresso Receipt
	espresso.coffeeAmmount = 20
	espresso.milkAmmount = 75
	espresso.moneyAmmount = 7
	espresso.waterAmmount = 350

	return espresso
}

func Cappucino() Receipt {
	//And for a cappuccino, the coffee machine needs 200 ml of water, 100 ml of milk, and 12 g of coffee beans. It costs $6.
	var espresso Receipt
	espresso.coffeeAmmount = 12
	espresso.milkAmmount = 100
	espresso.moneyAmmount = 6
	espresso.waterAmmount = 200

	return espresso
}

func status() {
	fmt.Println("The coffee machine has:")
	fmt.Println(available.waterAvailable, " ml of water")
	fmt.Println(available.milkAvailable, " ml of milk")
	fmt.Println(available.coffeeAvailable, " g of coffee beans")
	fmt.Println(available.cupsAvailable, " disposable cups")
	fmt.Printf("$%d of money\n", available.moneyAvailable)
	fmt.Println("")
}

func buyCoffee(available *Available) {

	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	var coffeeType int

	fmt.Scanln(&coffeeType)
	var coffee Receipt

	switch coffeeType {
	case 1:
		coffee = Espresso()
	case 2:
		coffee = Latte()
	case 3:
		coffee = Cappucino()
	default:
		break
	}

	canWeGetMoney := coffee.moneyAmmount != 0

	if canWeGetMoney && canWeGetMoneyFromUser(available, coffee) {
		available.moneyAvailable += coffee.moneyAmmount
		available.coffeeAvailable -= coffee.coffeeAmmount
		available.milkAvailable -= coffee.milkAmmount
		available.waterAvailable -= coffee.waterAmmount
		available.cupsAvailable -= 1
	}
}

func canWeGetMoneyFromUser(available *Available, receipt Receipt) bool {

	result := true

	if available.coffeeAvailable < receipt.coffeeAmmount {
		fmt.Println("Sorry, not enough coffee!")
		result = false
	} else if available.waterAvailable < receipt.waterAmmount {
		fmt.Println("Sorry, not enough water!")
		result = false
	} else if available.milkAvailable < receipt.milkAmmount {
		fmt.Println("Sorry, not enough milk!")
		result = false
	} else if available.cupsAvailable < 1 {
		fmt.Println("Sorry, not enough cups!")
		result = false
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
	}

	return result
}

func initializeCoffeeMachine() {
	// At the moment, the coffee machine has $550, 400 ml of water, 540 ml of milk, 120 g of coffee beans, and 9 disposable cups.
	available.moneyAvailable = 550
	available.waterAvailable = 400
	available.milkAvailable = 540
	available.coffeeAvailable = 120
	available.cupsAvailable = 9
}

func main() {
	initializeCoffeeMachine()

	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		var operation string
		fmt.Scanln(&operation)

		switch operation {
		case "buy":
			buyCoffee(&available)
		case "fill":
			fillCoffeeMachine()
		case "take":
			takeMoneyFromCoffeeMachine()
		case "remaining":
			status()
		case "exit":
			return
		default:
			fmt.Println("Incorrect command")
		}
	}
}

func fillCoffeeMachine() {

	fmt.Println("Write how many ml of water you want to add:")
	var addedWater int
	fmt.Scanln(&addedWater)
	available.waterAvailable += addedWater

	fmt.Println("Write how many ml of milk you want to add:")
	var addedMilk int
	fmt.Scanln(&addedMilk)
	available.milkAvailable += addedMilk

	fmt.Println("Write how many grams of coffee beans you want to add:")
	var addedCoffee int
	fmt.Scanln(&addedCoffee)
	available.coffeeAvailable += addedCoffee

	fmt.Println("Write how many disposable cups you want to add:")
	var addedCups int
	fmt.Scanln(&addedCups)
	available.cupsAvailable += addedCups
}

func takeMoneyFromCoffeeMachine() {
	fmt.Println("I gave you $", available.moneyAvailable)
	fmt.Println("")
	available.moneyAvailable -= available.moneyAvailable
}
