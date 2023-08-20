package main

import "fmt"

func main() {
	var title string
	var copies int
	var edition int
	var specialOffer bool
	var discountPercentage float64
	title = "For the Love of Go"
	copies = 99
	edition = 4
	specialOffer = true
	discountPercentage = 10
	fmt.Println(title)
	fmt.Println(copies, edition, specialOffer, discountPercentage)
}
