package main

import (
	"bufio"
	_ "bufio"
	. "fmt"
	"os"
	_ "os"
	"sort"
	"strings"
	"unicode"
	"regexp"
)

func main() {
	//questions()
	//first()
	//second()
	//third()
	//four()
}

func four() {
	m:= make(map[string]int)
	data := os.Args[1:]
	s := ""
	for l := range data{
		s+=data[l]
	}
	for i := 0; i <= len(s); i++ {
		for j := i+2; j < len(s); j++ {
			if s[i] == s[j]{
				m[s[i:j+1]] +=1
				//Println(m[s[i:j+1]],s[i:j+1])
			}
		}
	}
	for k := range m{
		Printf("%s -> Occorrenze: %d\n",strings.TrimSpace(strings.ReplaceAll(k,""," ")),m[k])
	}
}

func third() {
	mappa := ContaRipetizioni(SeparaParole(filter.ReplaceAllString(LeggiTesto()," ")))
	delete(mappa,"")
	Printf("Parole distinte: %d\n",len(mappa))
	for s := range mappa{
		Printf("%s: %d\n",s,mappa[s])
	}
}

func SeparaParole(s string) []string{
	return strings.Split(s," ")
}

var filter = regexp.MustCompile(`[^a-zA-Z]+`)
func ContaRipetizioni(s []string)  map[string]int{
	m := make(map[string]int)
	for _,k := range s{
		m[k] +=1
	}
	return m
}

func second() {
	s:= LeggiTesto()
	mappa := Occorrenze(s)
	chiavi := []string{}
	for runa := range mappa {
		chiavi = append(chiavi,string(runa))
	}
	sort.Strings(chiavi)
	for _,b := range chiavi{
		runa := []rune(b)[0]
		Printf("%s: %s\n",string(runa),strings.Repeat("*",mappa[runa]))
	}
}

func first(){
	s:= LeggiTesto()
	mappa := Occorrenze(s)
	Println("Istogramma:")
	for runa := range mappa{
		Printf("%s: %s\n",string(runa),strings.Repeat("*",mappa[runa]))
	}
}

func LeggiTesto() (s string){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		s+= scanner.Text() + "\n"
	}
	return
}

func Occorrenze(s string) map[rune]int{
	m:= make(map[rune]int)
	for _,runa := range s{
		if unicode.IsLetter(rune(runa)){
			m[runa] +=1
		}
	}
	return m
}

func questions() {
	// 2) f presente con valore 0
	//	elementi in mappa:
	//	a  10 \n	b -5  \n	d = 5

	// 3) C = 0  \n	B = -5 \n elementi in mappa:   key: A val: 110 	\n key:D val: 5
	// 4)   [nome],[num. telefonico]
	// 5)    [nome] , [tel1, tel2]
	// 6)
}

