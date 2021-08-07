package main

import (
	"errors"
	"fmt"
	"math"
)

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func avg1(x, y float64) float64 {
	return (x + y) / 2
}

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentcange := (currentPrice / prevPrice) * 100
	return change, percentcange
}

func getStockPriceChangeErrorHandling(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("previous price cannot be zero")
		return 0, 0, err
	}

	change := currentPrice - prevPrice
	percentcange := (currentPrice / prevPrice) * 100
	return change, percentcange, nil
}

func getStockPriceChangeNamedReturnValues(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (currentPrice / prevPrice) * 100
	return change, percentChange
}

func getStockPriceChangeNakedReturn(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (currentPrice / prevPrice) * 100
	return
}

func main() {
	x := 5.75
	y := 6.25

	result := avg(x, y)
	result1 := avg1(x, y)

	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)
	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result1)

	prevStockPrice := 75000.0
	currentStockPrice := 100000.0

	change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	}

	prevStockPriceError := 0.0
	currentStockPriceError := 100000.0

	change, percentChange, err := getStockPriceChangeErrorHandling(prevStockPriceError, currentStockPriceError)

	if err != nil {
		fmt.Println("Sorry! There was an error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
		} else {
			fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
		}
	}

	change1, percentChange1 := getStockPriceChangeNamedReturnValues(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change1), math.Abs(percentChange1))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change1, percentChange1)
	}

	change2, percentChange2 := getStockPriceChangeNakedReturn(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change2), math.Abs(percentChange2))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change2, percentChange2)
	}

	change3, _ := getStockPriceChangeNakedReturn(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change3))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change3)
	}
}
