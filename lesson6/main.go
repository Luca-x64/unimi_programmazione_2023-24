package main

import (
	. "fmt"
	"strconv"
	"strings"
	"unicode"
	"math"
)
type frazione struct{
	numeratore,denominatore int
}
func main() {
	//first()
	//repeat()
	//concat()
	//maiusc()
	//sumStrings()
	//lastSurname()
	//hiddenNumber()
	//frazioni()
	//minTermini()
	//DecToBaseX()
	//BaseXToHex()
}

func convertFromBaseXTo10(n string, b int) (sum int,r bool){
	r = true
	if b < 2 || b > 16{
		return 0,false
	}
	for i := len(n)-1; i >= 0; i-- {
		t := strings.Index(hex,string(n[i]))
		sum += t * int(math.Pow(float64(b),float64(len(n)-i-1)))
	}
	return sum,r	
}

func BaseXToHex() {
	var n string
	var b int
	Print("Enter num and base (letters only upper case):")
	Scan(&n,&b)
	a,_ := convertFromBaseXTo10(n,b)
	Print(a)
}

const hex string = "0123456789ABCDEF"
func convert10ToBaseX(n, b uint) (string, bool) {
	if b < 2 || b > 16{
		return "",false
	}
	var s string
	for n%b != 0 {
		s = string(hex[int(n%b)]) + s
		n/=b
	}
	return s,true
}

func DecToBaseX() {
	var n10,b uint
	Print("Enter num (b10) and new Base:")
	Scan(&n10,&b)
	s,_ := convert10ToBaseX(n10,b)
	Println(s)
}

func minTermini() {
	Println("Enter n,d: ")
	var n,d int
	Scan(&n,&d)
	if d != 0{
		f := NuovaFrazione(n,d)
		StampaFraz(Riduci(f))
	}
}

func Riduci(f *frazione) *frazione {
	for i := 2; i <= f.numeratore; i++ {
		if f.numeratore % i == 0 && f.denominatore % i == 0{
			f.numeratore /=i
			f.denominatore /=i
			i=1
		}
	}
	return f
}


func frazioni() {
	Println("Enter n,d: ")
	var n,d int
	Scan(&n,&d)
	f := NuovaFrazione(n,d)
	StampaFraz(f)
}

func StampaFraz(f *frazione) {
	Printf("Frazione: %d/%d",f.numeratore,f.denominatore)
}

func NuovaFrazione(num,denom int) *frazione {
	var f frazione = frazione{num,denom}
	return &f
}


func NumeroNascosto(testo string) (int, error){
	var num string 
	for _,e := range testo{
		if unicode.IsDigit(e){
			num += string(e)
		}
	}
	if len(num) == 0 {
		return 0,nil
	}
	return strconv.Atoi(num)
}


func hiddenNumber() {
	var s string
	Print("s: ")
	Scan(&s)
	n , _ := NumeroNascosto(s)
	if n == 0 {
		Println("Nessun numero nascosto!")
	} else {
	Println(n ," * 2 = ",n*2)
	}
}

func lastSurname() {
	var input,last string = "",""
	Println("Enter s:")
	Scan(&input)
	for input != "0" {
		if input > last{
			last = input
		}
		Print("Enter s:")
		Scan(&input)
	}
	Println(last)
}

func sumStrings() {
	var s string
	var sum int
	Println("Enter num:")
	Scan(&s)
	i,err := strconv.Atoi(s)
	for  err == nil {
		sum +=i
		Println("Enter num:")
		Scan(&s)
		i,err = strconv.Atoi(s)
	}
	Println("sum: ",sum)
}

func maiusc() {
	var s string
	Println("Enter s:")
	Scan(&s)
	for s != "0" {
		Println("maiusc:",strings.ToUpper(s))
		Print("Enter s:")
		Scan(&s)
	}
}

func concat() {
	var s,total string
	Println("Enter s:")
	Scan(&s)
	for s != "0" {
		total += s+" "
		Print("Enter s:")
		Scan(&s)
	}
	Println(total)
}

func repeat() {
	var s string
	Println("enter string:")
	Scan(&s)
	var n int
	Print("Enter n:")	
	Scan(&n)
	for i := 0; i < n; i++ {
		Print(s)
		if i< n-1{
			Print("-")
		}
	}

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
