package main

import (
	. "fmt"
	"strconv"
	"strings"
)

func main() {
	first()
	second()

}

func second() {
	var s string
	Println("enter:")
	Scan(&s)
	splt := strings.Split(" ")

}

func first() {
	var s string
	Println("enter n:")
	Scan(&s)
	Println(s)
	a, err := strconv.Atoi(s)
	Println(a, err)
	if a < 0 || err == nil {
		for i := 1; i <= a; i++ {
			Println(strings.Repeat("*", i))
		}

	} else {
		Println("error")
	}
}
