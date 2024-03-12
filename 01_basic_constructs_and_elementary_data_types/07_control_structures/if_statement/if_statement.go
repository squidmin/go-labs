package if_statement

import "fmt"

func demoIfStatement() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
