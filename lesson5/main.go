package main

import (
	. "fmt"
	. "strings"
	"unicode"

	_ "math"
)

func main(){
	// Domande
	// 1] > 5, 2.0 , 10
	// 2] bool
	// 3] c = a+b ; 
	// 4] -126 
	//      var a,b int8 = 30,100
	//      var somma int8 = a+b
	//      Println(somma)
	// 5] 15, 21, 22 
	// 6] Ne uno, ne due\n Ne uno, ne due\n 
	// 7] Ottimo\nInsuffieciente!\n
	// 8] 20
	// 9] 45
	// 10] 3) Minore di 5\n * (3)
	//	   100) Minore di 4\n *4 	Minore di 10 *5

//  [ Esercizi: ]
	//tabelline()
	//Scacchiera()
	//prime()
	//primiGemelli()
	//ternaPitagorica()
	//carte()
	//spaziatura_caratteri()
	//stringa_alternata()
	//analisiLettere()
	//menu()
}

func menu() {
	const patatine, hamburger, cocaCola, consegna int = 2,5,2,2
	var np,nh,nc int // number of items
	menu := "1. patatine 2€\n"+
	"2. hamburger 5€\n"+
	"3. cocacola 2€\n"+
	"0. termina"
	Println("Cosa vuoi ordinare?")
	Println(menu)
	var choise int
	Scan(&choise)
	for ; choise != 0;{
		switch choise{
		case 1:
			Print("Patatine? ottimo, quante? ")
			Scan(&np)	
		
		case 2:
			Print("Hamburger? ottimo, quante? ")
			Scan(&nh)	
		
		case 3:
			Print("cocacola? ottimo, quante? ")
			Scan(&nc)	
		}
		Println(menu)
		Scan(&choise)
		
	}
	var total int = patatine *np + hamburger * nh + cocaCola*nc + consegna
	Println("Sono, ",total-2," euro +2 di consegna. Totale: ",total)
}

func analisiLettere() {
	var s string
	var v,c,vm,cm int = 0,0,0,0
	Scan(&s)
	for _,char := range s{
		if unicode.IsLetter(char){

			if unicode.IsUpper(char){
				if èVocale(unicode.ToLower(char)){
					vm++
				}else{
					cm++
				}
			}else{
				if èVocale(char){
					v++
				}else{
					c++
				}
			}


		}
	}
	print(vm,cm,v,c)

}
func èVocale(l rune) bool{
	switch l{
	case 97,101,105,111,117: return true
	default: return false
	}
}

func stringa_alternata() {
	var s,l string
	Scan(&s,&l)
	if len(s) > len(l){
		l += Repeat("-",len(s)-len(l))
	}else if len(s) < len(l){
		s += Repeat("-",len(l)-len(s))
	}
	Println(StringheAlternate(s,l))
}
func StringheAlternate(s1, s2 string) (risultato string){

	for i := 0; i < len(s1); i++ {
	risultato+=string(s1[i])+string(s2[i])		
	}
	return
}

func spaziatura_caratteri() {
	var s string
	Print("Enter string: ")
	Scan(&s)
	for _,c := range s{
		Print(string(c)," ")
	}
}

func carte(){
	var start rune = 127153
	for i := start; i < start+10; i++ {
		Println("Simbolo: ", string(i)," Codice numero in base 10:",i)
	}
}


func ternaPitagorica() {
	var n int
	Print("soglia:")
	Scan(&n)
	if n <=0 {Println("La soglia non è positiva");return}
	TernePitagoriche(n)
		
		
	
}

func ETernaPitagoriga(a int, b int, c int) bool{
	return a*a + b*b == c*c
}

func TernePitagoriche(soglia int){
	for a := 3; a < soglia; a++ {
		for b := 3; b < soglia; b++ {
			for c := a+1; c < soglia; c++ {
				if ETernaPitagoriga(a,b,c){
					Printf("(%d, %d, %d)\n",a,b,c)
				}
			}
		}
}

}

func primiGemelli() {
	var n int
	Print("soglia:")
	Scan(&n)
	if n <=0 {Println("La soglia non è positiva");return}
	NumeriPrimiGemelli(n)

}
func NumeriPrimiGemelli(limite int){
	p,q :=1,3
	for i := 3; i < limite; i+=2 {

		if q < limite {
			p=q
			q=p+2
			if EPrimo(p) && EPrimo(q){
				Printf("(%d,%d) ",p,q)
			}
		}
	}
}

func prime() {

	var n int
	Print("soglia:")
	Scan(&n)
	if n <=0 {Println("La soglia non è positiva");return}
	NumeriPrimi(n)

}

func EPrimo(n int) bool{
	if n % 2 == 0 && n > 2{return false}
	for i := 3; i < n; i+=2 {
		if n%i == 0{
			return false
		}
	}
	return true
}

func NumeriPrimi(limite int) {
	for i := 2; i < limite; i++ {
		if EPrimo(i){
			Print(i," ")
		}
	}
}

func Scacchiera() {
	var n int
	Print("n:")
	Scan(&n)
	StampaScacchiera(n)
	
}
func StampaScacchiera(size int){
	if size <= 0 {return}
	var flag bool = true
	for i := 0; i < size; i++ {
		if flag{
			StampaRigaInizioAsterisco(size)
		}else{
			StampaRigaInizioPiù(size)
		}
		
		flag = !flag
		
	}


}

func StampaRigaInizioAsterisco(size int) {
	var s string = "*+"
	Println(Repeat(s,size/2 + size%2)[:size])
}
func StampaRigaInizioPiù(size int){
	var s string = "+*"
	Println(Repeat(s,size/2 + size%2)[:size])
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
		Print("Tabellina del",n,": ")
		for i := 1; i <= 10; i++ {
			Print(i*n," ")
		}
		return true
	}else{
		Println("Programma terminato")
		return false
	}
}