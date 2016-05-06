package main

import "fmt"

func Move(n, from, to, temp int) {
	if n <= 0 {
		return
	}
	Move(n-1, from, temp, to)
	fmt.Println(from, "->", to)
	Move(n-1, temp, to, from)
}

func Hanoi(n int) {
	fmt.Println("Number of disks:", n)
	Move(n, 1, 2, 3)
}

func main() {
	Hanoi(3)
}
