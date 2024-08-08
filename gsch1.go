package main

func main() {
	s := "abc"
	println(&s)
	s, y := "def", 10
	println(&s, y)

	{
		s, z := 1000, 30

		println(&s, z)
	}
}
