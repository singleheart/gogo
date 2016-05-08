package array

import (
	"fmt"

	"github.com/singleheart/gogo/hangul"
)

func Example() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func getPostposition(s string) string {
	if hangul.HasConsonantSuffix(s) {
		return "은"
	} else {
		return "는"
	}
}

func Example_ConsonantSuffix() {
	fruits := [...]string{"사과", "바나나", "토마토", "자몽"}
	for _, fruit := range fruits {
		fmt.Printf("%s%s 맛있다.\n", fruit, getPostposition(fruit))
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 자몽은 맛있다.
}
