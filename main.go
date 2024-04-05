//Made by yacine aridhi with loveee

package main

import (
	"fmt"
)

func main() {
	// a loop generates numbers until 100 is reached and is stored in a variable called i
	for i := 0; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			// check if the i is a common multiple of 3 and 5 print fizz buzz
			fmt.Println(i, " Fizz Buzz")
		} else if i%3 == 0 {
			// check if the i is a multiple of 3 print fizz
			fmt.Println(i, " Fizz")
		} else if i%5 == 0 {
			// check if the i is a multiple of 5 print buzz
			fmt.Println(i, " Buzz")
		} else {
			// print numbers i that are neither multiples of 3 nor 5
			fmt.Println(i)
		}

	}
}
