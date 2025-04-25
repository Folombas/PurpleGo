package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
	
	


	fmt.Println("__ Калькулятор степени ожирения __")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У ваc дефицит массы тела")	
	case IMT < 25:
		fmt.Println("У ваc нормальный вес")
	case IMT < 30: 
		fmt.Println("У ваc избыточный вес")
	default:
		fmt.Println("У ваc степень ожирения")	
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
