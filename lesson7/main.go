package main

import (
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//questions()
	//first()
	//Inverso()
	checkSequence()
}

func checkSequence() {
	args := os.Args[1:]
	var flag bool = check(args)
	Print(flag)
}

func check(s []string) bool {
	var save []int 
	for i := 0; i < len(s); i++ {
		n,err := strconv.Atoi(s[i])
		if err != nil{ // check non number
			return false
		}else{
			save = append(save,n)
			pos := len(save) 
			if pos % 2 != 0{
				if save[pos] >= save[pos-1]{ // second check	
					return false
				}
			}else {
				if pos > 1 && save[pos] <= save[pos-1]{ // third check
					return false
				}
			}
		}
	}
	return true
}

func Inverso() {
	var n int
	Print("Enter n:")
	Scan(&n)
	var numeri []int
	for i := 0; i < n; i++ {
		var t int
		Scan(&t)
		numeri = append(numeri,t)
	}
	for i := len(numeri)-1; i >= 0; i-- {
		Print(numeri[i], " ",)
	}
	




	
}

func first() {
	Occorrenze("Ciaocomestai") // ?
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func Occorrenze(s string) (count [26]int) {
	s = strings.ToLower(s)
	for len(s) != 0 {
		var pos int = strings.Index(alphabet, string(s[0]))
		sub := strings.ReplaceAll(s, string(s[0]), "")
		count[pos] = len(s) - len(sub)
		s = sub
	}
	for i := 0; i < len(count); i++ {
		Print(string(alphabet[i]), " ", (count[i]), " ")
	}
	return
}
func questions() {
	// 'abc'
	// 'helloworld'
	// '0\n1\n2\n3\n4\n5\n
	// [1 2 3 4 5 6]
	// [0 1 2 3 4]
	// [0 2 4 5 8 10]]
	// [0 1 0 0 5 6]
	//  	1 2 3 4 5
	//  	2 4 6 8 10
	//  	2 4 6 8 10
	// [10 9 8 7 6 5 4 3 2 1]  \n [5 4 3 2 1 0 0 0 0 0]
	// TIPO os.Args: array \n os.Args[0:4] : TIPO {typovalore i-esimo}  - VALORE = {valore stringa iesimo}\n"
	//												{int, string, float, string, bool}
}
