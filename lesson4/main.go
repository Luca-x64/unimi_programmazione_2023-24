package main

import . "fmt"
import "strings"

func main() {
	//quadrato()
	//quadratoAlterno()
	//quadratoTriploAlterno()
	//countString()
	//swapCase()
	palindroma()
}

func palindroma() {
	var s string
	var palindroma bool = true
	Scan(&s)
	
	// for i,_ := range s{
	// 	if (i >= len(s)/2){
	// 		break
	// 	}
	// 	if strings.Compare(string(s[i]), string(s[len(s)-1-i])) != 0{
	// 		palindroma=false
	// 	}
	// }
	//Println("Palindroma? ",palindroma)
	
	Println("Palindroma? ",palindroma)
}

func swapCase() {
	var s string
	Scan(&s)
	Println(strings.ToUpper(s),"\n",strings.ToLower(s))
}

func countString() {
	var s,sTemp string = "",""
	var temp rune
	var mai, min, con, voc int
	Scan(&s)
	for i := 0; i < len(s); i++ {
		temp = rune(s[i])
		sTemp = string(temp)
		if (temp >= 65 && temp <= 90) || (temp >= 97 && temp <= 122) {
			if sTemp == strings.ToUpper(sTemp) {
				mai++
			} else {
				min++
			}
			sTemp = strings.ToLower(sTemp)
			if (sTemp[0] == 97) || (sTemp[0] == 101) || (sTemp[0] == 105) || (sTemp[0] == 111) || (sTemp[0] == 117) {
				voc++
			} else {
				con++
			}
		}
	}
	Print("maiuscole: ",mai,"\nmin:",min,"\nconsonanti:",con,"\nvocali:",voc,"\n")
}

func quadratoTriploAlterno() {
	var n int
	simbol := -1
	Print("enter n: ")
	Scan(&n)
	for i := 0; i < n; i++ {
		simbol = (simbol + 1) % 3
		for i := 0; i < n; i++ {
			if simbol == 0 {
				Print("*")
			} else if simbol == 1 {
				Print("+")
			} else {
				Print("o")
			}
		}
		Println()
	}
}

func quadratoAlterno() {
	var n int
	var flag bool = false
	Print("enter n: ")
	Scan(&n)
	for i := 0; i < n; i++ {
		flag = !flag
		for i := 0; i < n; i++ {
			if flag {
				Print("*")
			} else {
				Print("+")
			}
		}
		Println()
	}
}

func quadrato() {
	var n int
	Print("enter n: ")
	Scan(&n)
	for i := 0; i < n; i++ {
		for i := 0; i < n; i++ {
			Print("*")
		}
		Println()
	}
}
