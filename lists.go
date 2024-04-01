package main

import "fmt"

func main() {
	//var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	//productNames[2] = "A carpet"
	//fmt.Println(productNames)
	//fmt.Println(prices[2])

	featuredPrice := prices[1:]
	fmt.Println(featuredPrice)
	highlitedPrices := featuredPrice[:1]
	fmt.Println(highlitedPrices)

}
