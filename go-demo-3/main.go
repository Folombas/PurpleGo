package main

import "fmt"

func main() {
	m := map[string]string{
		"Google": "google.com",
	}
	fmt.Println(m)
	fmt.Println(m["Google"])
	m["Google"] = "https://google.com"
	fmt.Println(m)
	m["Yandex"] = "https://yandex.ru"
	m["Rutube"] = "https://rutube.ru"
	fmt.Println(m)
	delete(m, "Yandex")
	delete(m, "Y")
	fmt.Println(m["Y"])
	fmt.Println(m)
}