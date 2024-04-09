package main

import "fmt"

func main() {

	var t X
	t.a = 10
	t.l = []int{1, 2, 3, 4, 5}

	xxx(t)

}

type X struct {
	a int
	l []int
}

func xxx(x X) {
	fmt.Println("x show => ", x)
}
