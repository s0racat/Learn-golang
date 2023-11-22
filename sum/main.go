package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int = 0
	for i := 1; i <= 2; i++ {
		// fmt.Println(os.Args[i])
		num, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	}
	fmt.Println(sum)
}
