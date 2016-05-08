package strings

import (
	"fmt"
	"strconv"
)

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	s := []byte("가나다")
	s[2]++
	fmt.Println(string(s))
	// Output:
	// 각나다
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_strToNumber() {
	var i int
	var k1 int64
	var k2 int64
	var f float64

	i, _ = strconv.Atoi("350")
	k1, _ = strconv.ParseInt("cc7fdd", 16, 32)
	k2, _ = strconv.ParseInt("0xcc7fdd", 0, 32)
	f, _ = strconv.ParseFloat("3.14", 64)

	fmt.Println(i)
	fmt.Println(k1)
	fmt.Println(k2)
	fmt.Println(f)
	fmt.Println(strconv.Itoa(340))
	fmt.Println(strconv.FormatInt(13402077, 16))

	var num int
	fmt.Sscanf("57", "%d", &num)
	fmt.Println(num)

	var s string
	s = fmt.Sprint(3.14)
	fmt.Println(s)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(s)

	// Output:
	// 350
	// 13402077
	// 13402077
	// 3.14
	// 340
	// cc7fdd
	// 57
	// 3.14
	// cc7fdd
}
