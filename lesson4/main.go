package main

import (
	. "fmt"
	"strings"
)

func main() {
	//quadrato()
	//quadratoAlterno()
	//quadratoTriploAlterno()
	//countString()
	//swapCase()
	//palindroma()
	triangle()
}

func triangle() {
	var n int
	Print("n:")
	Scan(&n)
	top(n)
	bottom(n)
}

func stampa(n int) {
	for i := 0; i < n; i++ {
		Print("*")
	}
}

func top(n int) {
	for i := 0; i < n; i++ {
		if i == 0 {
			stampa(n)
		} else {
			for j := 0; j < n; j++ {
				if j == n-1 || j == i {
					Print("*")
				} else {
					Print(" ")
				}
			}
		}
		Println()
	}
}
func bottom(n int) {
	for i := 0; i < n; i++ {
		if i == n-1 {
			stampa(n)
		} else {
			for j := 0; j < n; j++ {
				if j == n-1 || j == i {
					Print("*")
				} else {
					Print(" ")
				}
			}
		}
		Println()
	}
}

func palindroma() {
	var s string
	var palindroma bool = true
	Scan(&s)
	for i := 0; i < len([]rune(s))/2; i++ {
		if []rune(s)[i] != []rune(s)[len([]rune(s))-1-i] {
			palindroma = false
		}
	}
	Println("Palindroma? ", palindroma)
}

func swapCase() {
	var s string
	Scan(&s)
	Println(strings.ToUpper(s), "\n", strings.ToLower(s))
}

func countString() {
	var s, sTemp string = "", ""
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
	Print("maiuscole: ", mai, "\nmin:", min, "\nconsonanti:", con, "\nvocali:", voc, "\n")
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
