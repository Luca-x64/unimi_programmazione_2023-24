package main

import (
	. "fmt"
	"strings"
)

func main() {
	//questions()
	first()
}

func first() {
	Occorrenze("Ciaocomestai") // ?
}
const alphabet = "abcdefghijklmnopqrstuvwxyz"
func Occorrenze(s string) (count [26]int){
	for _ , v := range s{
		var pos int = strings.Index(alphabet,string(v))
		count[pos] = len(s)-1 - len(strings.ReplaceAll(s,string(s[0]),""))
	}
	Print(count)
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
