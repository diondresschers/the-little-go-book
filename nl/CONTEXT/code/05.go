package main

import "fmt"

func main() {
	kracht := 1000
	fmt.Printf("de standaard kracht is %d\n", kracht)

	naam, kracht := "Goku", 9000
	fmt.Printf("%s's kracht is meer dan %d\n", naam, kracht)
}
