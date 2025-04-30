package main

import (
	"fmt"
)

func main() {
	login := promptData("Введите Ваш логин")
	password := promptData("Введите Ваш пароль")
	url := promptData("Введите URL")
	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}


