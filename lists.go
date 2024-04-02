package main

import "fmt"

func main() {

	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	//fmt.Println(prices[:1])
	prices = append(prices, 5.99, 12.99, 29.99, 100.0)
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)

	discontPrices := []float64{101.99, 80.99, 20.5}
	prices = append(prices, discontPrices...)
	fmt.Println(prices)
	/*
		var productNames [4]string = [4]string{"A book"}
		prices := [4]float64{10.99, 9.99, 45.99, 20.0}
		fmt.Println(prices)
		productNames[2] = "A carpet"
		fmt.Println(productNames)
		fmt.Println(prices[2])
		featuredPrice := prices[1:]
		fmt.Println(featuredPrice)
		highlitedPrices := featuredPrice[:1]
		featuredPrice[0] = 199.99
		fmt.Println(highlitedPrices)
		fmt.Println(len(featuredPrice), cap(featuredPrice))
		fmt.Println(prices)
	*/
}
