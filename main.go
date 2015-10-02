package main

import "fmt"

//Using Closure: method 1

func generator() func() int {
	i := int(0)
	return func() (ret int) {
		ret = i
		i++
		if i%2 == 0 {
			fmt.Println(i/2, "True")
		} else {
			fmt.Println(i/2, "False")
		}
		return
	}
}
func main() {
	number := generator()
	number() // 0
	number() // 2
	number() // 4
	number() // 4
}

//Using Closure: method 2

func main() {
	x := 0
	increment := func() int {
		x++
		if x%2 == .0 {
			fmt.Println(x/2, "True")
		} else {
			fmt.Println(x/2, "false")
		}
		return x
	}
	increment()
	increment()
	increment()
	increment()
	increment()
	increment()
	increment()
}
