package main

import "fmt"

func main() {
	jam := 10
	menit := 5
	detik := jam*3600 + menit*60

	fmt.Println(jam, "jam", "dan", menit, "menit", "adalah", detik, "detik")
}
