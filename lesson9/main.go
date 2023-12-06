package main

import (
	. "fmt"
	"math/rand"
	"strconv"
	"strings"
	_ "testing"
)

func main() {
	//test()
	//exampleOddOrEven()
	//carteDaGioco()
	//subStringPalindrome()
	balancedBrackets()
}

func balancedBrackets() {
	var s string
	Print("Enter s:")
	Scan(&s)
	if isBalanced(s){
		Println("bilanciata")
	}else{
		Println("non bilanciata")
	}
	//parte 2


}
func isBalanced(s string) bool{
	index := strings.Index(s,"()")
	copy := s
	for index != -1{
		copy =strings.ReplaceAll(copy,"()","")
		index = strings.Index(copy,"()")
	}
	return len(copy) == 0
}

func subStringPalindrome() {
	var s string
	Print("Enter s:")
	Scan(&s)
	ln := len(s)
	for i := 0; i < ln; i++ {
		for j := 2; j < ln-i+1; j++ {
			sub := s[i:i+j]
			if isPalindrome(sub){
				Println(sub)
			}
		}
	}

}
func isPalindrome(s string)  bool{
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i]{
			return false
		}
	}
	return true
}

type Carta struct {
	Seme    string
	Simbolo string
}
type Mazzo struct {
	carte  []Carta
	nCarte int
}

func carteDaGioco() {
	var mazzo Mazzo = creaMazzo()
	Println("Mazzo creato:", mazzo)
	mazzo = shuffle(mazzo)
	Println("\nMazzo mescolato:", mazzo)
	Println()
	m := mazzo
	for i := 0; i < 41; i++ {

		c, mv, err := Preleva(m)
		if err != nil {
			Println(err)
		} else {
			Println("Prelevata:", c, "totale carte: ", m.nCarte)
			m = mv
		}

	}

	for i := 0; i < 41; i++ {
		rnd := rand.Intn(40)
		var c Carta = mazzo.carte[rnd]
		mv, err := Riponi(m, c)
		if err != nil {
			Println(err)
		} else {
			Println("Inserita", c, "totale carte:", m.nCarte)
			m = mv
		}

	}
}
func creaCarta(seme, simbolo string) (c Carta) {
	c.Seme = seme
	c.Simbolo = simbolo
	return
}

func creaMazzo() (m Mazzo) {
	semi := [4]string{"♦", "♠", "♥", "♣"}
	simboli := [10]string{"asso", "2", "3", "4", "5", "6", "7", "fante", "donna", "re"}
	for seme := range semi {
		for simbolo := range simboli {
			// pos,_ := strconv.Atoi(strconv.Itoa(seme)+strconv.Itoa(simbolo))
			m.carte = append(m.carte, creaCarta(semi[seme], simboli[simbolo]))
			m.nCarte = m.nCarte + 1
		}
	}

	return
}
func shuffle(m Mazzo) Mazzo {
	for i := 0; i < m.nCarte; i++ {
		rnd := rand.Intn(i + 1)
		m.carte[i], m.carte[rnd] = m.carte[rnd], m.carte[i]
	}
	return m

}
func Preleva(m Mazzo) (c Carta, mm Mazzo, e error) {
	if m.nCarte > 0 {
		c = m.carte[0]
		m.nCarte = m.nCarte - 1
		mm = Mazzo{m.carte[1:], m.nCarte}
		e = nil
	} else {
		e = Errorf("Mazzo vuoto")
	}
	return
}
func Riponi(m Mazzo, c Carta) (new Mazzo, e error) {
	new = m
	if m.nCarte < 40 {
		new.nCarte = new.nCarte + 1
		new.carte = append(m.carte, c)
		e = nil
	} else {
		e = Errorf("Mazzo Pieno")
	}
	return
}

func exampleOddOrEven() {
	var input string
	Scan(&input)
	n, _ := strconv.Atoi(input)
	Println(oddOrEven(n))
}

func oddOrEven(n int) bool {
	return n%2 == 0
}
func test() {
	var input string
	Scan(&input)
	n, err := strconv.Atoi(input)
	if err == nil {
		Println("il valore inserito è un numero reale")
		if n == 0 {
			Println("il valore inserito è un numero intero nullo")
		} else if n < 0 {
			Println("il valore inserito è un numero intero negativo")
		} else if n == 1 {
			Println("il valore inserito è 1 (1 non è un numero primo)")
		} else {
			if isprime(n) {
				Println("il valore inserito è un numero intero positivo primo")
			} else {
				Println("il valore inserito è un numero intero positivo non primo")
			}
		}

	} else {
		Println("il valore inserito non è un numero")
	}
}

func isprime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
