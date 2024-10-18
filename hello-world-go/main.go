package main

import "fmt"

var age1 int

func hello(name string) {
	fmt.Println("Hello", name)
}

func variables() {
	age2 := 12
	fmt.Println(age1)
	fmt.Println(age2)
}

func sum(a, b int) int {
	return (a + b)
}

func swap(a, b int) (int, int) {
	return b, a
}

func loop(i int) {
	for e := 1; e <= i; e++ {
		if e%2 == 0 {
			fmt.Println(e)
		} else {
			fmt.Println("Odd")
		}
	}
}

func price(item string) int {
	switch item {
	case "apple", "tomatoes":
		return 10
	case "orange":
		return 5
	case "banana":
		return 1
	default:
		return 0
	}
}

func main() {
	hello("Anurag")
	hello("Negi")
	variables()
	fmt.Println("Sum is:", sum(4, 6))
	a := 10
	b := 12
	a, b = swap(a, b)
	fmt.Print("a:", a, ", b:", b, "\n")
	loop(10)
	fmt.Println(price("apple"))
	fmt.Println(price("orange"))
	fmt.Println(price("banana"))
	fmt.Println(price("tomatoes"))
	fmt.Println(price("carrot"))
}
