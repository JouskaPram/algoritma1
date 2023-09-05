package main

import "fmt"

func main() {
	bilangan := 101
	genap := bilangan % 2

	if genap == 0 {
		fmt.Println(bilangan, "bilangan genap")
	} else {
		fmt.Println(bilangan, "bilangan ganjil")
	}

}
