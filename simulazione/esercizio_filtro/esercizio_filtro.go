package main

import (
	. "fmt"
	"os"
	"strconv"
)

func main(){
	var input string = os.Args[1]
	var output string
	prec,_ := strconv.Atoi(string(input[0]))
	for i := 1; i < len(input); i++ {
		n,_ := strconv.Atoi(string(input[i]))
		if prec < n{
			output+=strconv.Itoa(prec)
		}
		prec = n
	}
	Println(output)
}