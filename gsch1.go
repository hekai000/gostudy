package main

func exp1() {
	s := "abc"
	println(&s)
	s, y := "def", 10
	println(&s, y)

	{
		s, z := 1000, 30

		println(&s, z)
	}
	a := []int{1, 2, 3}
	a[1] = 10
	println(a)
	b := make([]int, 3)
	b[1] = 10
	println(b)
	c := new([]int)

	*c = make([]int, 3)
	(*c)[1] = 10
	println(*c)
}
