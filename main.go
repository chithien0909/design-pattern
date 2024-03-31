package main

import "fmt"

func main() {
	adidasF, _ := GetProductFactory("apple")
	nikeF, _ := GetProductFactory("orange")

	appleJuice := adidasF.makeJuice()
	appleJam := adidasF.makeJam()

	orangeJuice := nikeF.makeJuice()
	orangeJam := nikeF.makeJam()

	printJuice(appleJuice)
	printJam(appleJam)

	printJuice(orangeJuice)
	printJam(orangeJam)

}

func printJuice(j IJuice) {
	fmt.Println("Juice Name: ", j.getName())
	fmt.Println("Juice Price: ", j.getPrice())
	fmt.Println("====================================")
}

func printJam(j IJam) {
	fmt.Println("Jam Name: ", j.getName())
	fmt.Println("Jam Price: ", j.getPrice())
	fmt.Println("====================================")
}

/**
Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.
Single Responsibility Principle. You can isolate concrete product classes from the client code.
Open/Closed Principle. You can introduce new variants of products without breaking existing client code.
Avoiding tight coupling between concrete products and client code.
*/
