package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int = 1
	num, _ := strconv.Atoi(os.Args[1])
	for i := 1; i <= num; i++ {
		sum *= i
	}
	fmt.Println(sum)
}
