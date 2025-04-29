package main

import "fmt"

func main() {
	login := promptData("Введите Ваш логин")
	password := promptData("Введите Ваш пароль")
	url := promptData("Введите URL")

	outputPassword(login, password, url)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}
