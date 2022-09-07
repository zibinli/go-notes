package main

import "fmt"

func main()  {
	count := 3
	count2 := 3
	// OuterLop:
		for i := 1; i < count; i++ {
			for j := 1; j < count2; j++ {
				if j == 2 {
					// continue OuterLop
				}
				fmt.Println(i * j)
			}
		}
}