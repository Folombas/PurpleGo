package main

import "fmt"

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавлять кажду в массив транзакций
// Вывести массив

func main() {
	tr1 := []int{2, 4, 6}
	tr2 := []int{8, 10, 12}
	tr1 = append(tr1, tr2...)
	fmt.Println(tr1)

	for index, value := range tr1 {
		fmt.Println(index, value)
	}


	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}