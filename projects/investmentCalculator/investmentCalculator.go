package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.5
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
