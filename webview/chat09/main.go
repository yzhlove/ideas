package main

import "fmt"

func main() {

	fmt.Println(calc(9))
	fmt.Println(calc(98))
	fmt.Println(calc(1234))
	fmt.Println(calc(12034))
	fmt.Println(calc(110023))
	fmt.Println(calc(11000))
	fmt.Println(calc(33445))
	fmt.Println(calc(26006))

}

func calc(n int) []int {
	if n < 10 {
		return []int{1, n - 1}
	}
	x := n / 2
	var numbers []int
	for x > 0 {
		numbers = append(numbers, x%10)
		x /= 10
	}

	var sub = make([]int, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		if i == len(numbers)-1 {
			if numbers[i]-1 <= 0 {
				sub = append(sub, 0)
			} else {
				sub = append(sub, numbers[i]-1)
			}
		} else {
			switch numbers[i] {
			case 0:
				sub = append(sub, 5)
			case 1:
				sub = append(sub, 1)
			default:
				sub = append(sub, numbers[i]/2)
			}
		}
	}

	fmt.Println("n = ", n)
	fmt.Println("numbers = ", numbers)
	fmt.Println("sub --> ", sub)
	fmt.Println("----------------------------")
	return numbers
}

/*
 1234

    4
 123
   3
 12
  2
  1
1
0

*/
