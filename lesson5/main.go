package main

import . "fmt"

func main(){
	// 1] > 5, 2.0 , 10
	// 2] bool
	// 3] c = a+b ; 
	// 4] -126 
	//      var a,b int8 = 30,100
	//      var somma int8 = a+b
	//      Println(somma)
	// 5]
	tabelline()
}

func tabelline() {
	var num int
	Print("num: ")
	Scan(&num)
	for ; Tabellina(num);{
		Println()
		Print("num: ")
		Scan(&num)
	}

	
}

func Tabellina(n int) bool{
	if n>= 1 && n <= 9{
		Print("Tabellina del")
		for i := 0; i < 10; i++ {
			Print(i*n," ")
		}
		return true
	}else{
		return false
	}
}