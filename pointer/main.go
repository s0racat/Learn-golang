package main

import "fmt"

func main() {
	hoge := "test"
	p := &hoge
	fmt.Println(*p)
}