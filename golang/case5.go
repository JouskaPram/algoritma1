package main

import "fmt"

func main() {
	bilangan1 := 10
	bilangan2 := 100
	bilangan3 := 40

	var himpunan = [...]int{bilangan1, bilangan2, bilangan3}
	min, max := findMinAndMax(himpunan)
	avarage := himpunan[0] + himpunan[1] + himpunan[2]/3
	fmt.Println("Angka Tertinggi", max)
	fmt.Println("Angka Terendah", min)
	fmt.Println("Rata Rata", avarage)
}
func findMinAndMax(himpunan [3]int) (min int, max int) {
	min = himpunan[0]
	max = himpunan[0]
	for _, value := range himpunan {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
