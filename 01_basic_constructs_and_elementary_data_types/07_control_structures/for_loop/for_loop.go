package for_loop

import "fmt"

func demoForLoop() {
	fmt.Println("First for loop:")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println()

	fmt.Println("Second for loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	fmt.Println("Third for loop")
	for i := 9; i > 0; i-- {
		fmt.Println(i)
	}
}
