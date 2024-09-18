package main

import "fmt"

func changeStr() {
	s := "Hello, world!"
	bs := []byte(s)
	bs[1] = 'B'
	fmt.Println(string(bs))

	u := "你好，世界！"
	bu := []rune(u)
	bu[1] = '话'
	fmt.Println(string(bu))

}
func visitbr() {
	s := "abc汉字"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()

	for _, r := range s {
		fmt.Printf("%c,", r)
	}
}

func main() {
	visitbr()
}

// Output:
// HBlo, world!
// 你话，世界！
