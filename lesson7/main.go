package main

import (
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//questions()
	//first()
	//Inverso()
	//checkSequence()
	//filterMultiply()
	//minMaxMean()
	//fattoriale()
}

func fattoriale() {
	n,_ := strconv.Atoi(os.Args[1])
	Println("Fattoriale",n,":",fact(n))
}

func fact(n int) (f []int){
	if n > 0 {
		f = append(f,1)
	}
	for i := 2; i <= n; i++ {
		f= append(f,f[len(f)-1]*i)
	}
	return f
}

func minMaxMean() {
	numeri := LeggiNumeri()
	if len(numeri) == 0{
		Println("Enter args")
		return
	}
	Printf("Minimo: %d\nMassimo: %d\n",min(numeri),max(numeri))
	Printf(Sprintf("Media: %%.%df\n",2),mean(numeri))
}

func min(sl []int) int{
	var min int = math.MaxInt32
	for _,v := range sl{
		if v < min{
			min = v
		}
	}
	return min
}

func max(sl []int) int{
	var max int = math.MinInt32
	for _,v := range sl{
		if v > max{
			max = v
		}
	}
	return max
}

func mean(sl []int) float64{
	var sum int = 0
	for _,v := range sl{
		sum +=v
	}
	return (float64(sum)/float64(len(sl)))
}

func LeggiNumeri() (numeri []int){
	args := os.Args[1:]
	if len(args) == 0 {
		return	
	}

	for _,v := range args {
		n, err := strconv.Atoi(v)
		if err == nil {
			numeri = append(numeri,n)
		}
	}
	return
}

func filterMultiply() {
	args := os.Args[1:]
	if len(args) == 0 {
		return	
	}
	var product int = 1
	for _,v := range args {
		n, err := strconv.Atoi(v)
		if err == nil {
			product*=n
		}
	}
	Printf("Product: %d\n",product)
}

func checkSequence() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"mamma","mia!"}
	}
	check(args)
}

func check(s []string) bool {
	var save []int 
	for i := 0; i < len(s); i++ {
		n,err := strconv.Atoi(s[i])
		if err != nil{ // check non number
			Printf("valore in posizione %d non valido",i)
			return false
		}else{
			save = append(save,n)
			pos := len(save) -1
			if pos % 2 != 0{
				if save[pos] >= save[pos-1]{ // second check	
					Printf("valore in posizione %d non valido",i)
					return false
				}
			}else {
				if pos > 1 && save[pos] <= save[pos-1]{ // third check
					Printf("valore in posizione %d non valido",i)
					return false
				}
			}
		}
	}
	Println("Sequenza valida")
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
