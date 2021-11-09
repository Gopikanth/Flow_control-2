package main

import "fmt"

/*func main() {
	for i := 0; i <= 100; i++ {
		if i%5 == 0 {
			fmt.Println(i)
		}
	}
} */

//the same code can be executed in efficient way as
func main() {
	for i := 0; i <= 100; i = i + 5 {
		fmt.Println(i)
	}
}

//the  first_above code  the loop is executed 100 times
//whereas in the above code the loop is executed only 15 times
