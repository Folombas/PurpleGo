package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	str := []rune("Привет!)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}


	login := promptData("Введите Ваш логин")
	password := promptData("Введите Ваш пароль")
	url := promptData("Введите URL")

	myAccount := account{
		password: password,
		url:      url,
		login:	login,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}
