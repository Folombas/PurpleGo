package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	files.WriteFile("Привет! Я файл!!", "file.txt")
	login := promptData("Введите Ваш логин")
	password := promptData("Введите Ваш пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	myAccount.OutputPassword()
	
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}


