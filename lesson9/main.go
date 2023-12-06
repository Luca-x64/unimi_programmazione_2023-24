package main

import (
	. "fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	_ "testing"
)

func main() {
	//test()
	//exampleOddOrEven()
	//carteDaGioco()
	//subStringPalindrome()
	//balancedBrackets()
	//SubstringGrowing()
	//triangleTartaglia()
}

func triangleTartaglia(){
	n,e := strconv.Atoi(os.Args[1])
	if e == nil{
		StampaTriangolo(generaTriangolo(n))
	}
}
func generaTriangolo(n int) (t [][]int){
	for i := 0; i < n+1; i++ {
		line := []int{}
		for j := 0; j <=i; j++ {
			if j == 0 || j == i{
				line = append(line,1)
			}else{
				line = append(line,t[i-1][j-1]+t[i-1][j])
			}
		}
		t = append(t,line)
	}
	return 
}
func StampaTriangolo (t [][]int){
	for i := 0; i < len(t); i++ {
		for j := 0; j <= i; j++ {
			Print(t[i][j]," ")
		}
		Println()
	}
}

func SubstringGrowing() {
	s := os.Args[1]
	_, e := strconv.Atoi(s)
	if e != nil {
		Print([]string{})
	}else{
		sequence := sottoStringhe(s)
		Println(sequence)
	}
	
	
}
func sottoStringhe(s string) []string{
	sequence := []string{}
	for i := 0; i < len(s); i++ {
		for j := 2; j < len(s)-i+1; j++ {
			t := s[i : i+j]
			if len(t) >= 2 && grow(t) {
				sequence = append(sequence,t)
			}
		}
	}
	return sequence
}
func grow(t string) bool {
	for i := 0; i < len(t)-1; i++ {
		a, _ := strconv.Atoi(string(t[i]))
		b, _ := strconv.Atoi(string(t[i+1]))
		if a >= b {
			return false
		}
	}
	return true

}
func balancedBrackets() {
	var s string = os.Args[1]
	ln := len(s)
	if strings.Index(s, "(") == -1 && strings.Index(s, ")") == -1 {
		Println("L'input fornito non aveva esclusivamente parentesi tonde.")
		return
	}
	if isBalanced(s) {
		Println("bilanciata")
	} else {
		Println("non bilanciata")
	}
	Print("---\nSottosequenze bilanciate:\n")

	//parte 2
	cnt := 1
	for i := 0; i < ln; i++ {
		for j := 2; j < ln-i+1; j++ {
			sub := s[i : i+j]
			if isBalanced(sub) {
				Println(cnt, sub)
				cnt += 1
			}
		}
	}
	if cnt == 1 {
		Println("Nessuna")
	}

}
func isBalanced(s string) bool {
	index := strings.Index(s, "()")
	copy := s
	for index != -1 {
		copy = strings.ReplaceAll(copy, "()", "")
		index = strings.Index(copy, "()")
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
			sub := s[i : i+j]
			if isPalindrome(sub) {
				Println(sub)
			}
		}
	}

}
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
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
