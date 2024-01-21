package main

import (
	. "fmt"
	"os"
	"strconv"
	"strings"
	"math")

func main() { // 1h, 02minuti
	var equation string = os.Args[1]
	var soglia,epsilon float64
	soglia ,_ = strconv.ParseFloat(os.Args[2],64)
	epsilon ,_ = strconv.ParseFloat(os.Args[3],64)
	//
	var a, b, c, right_c int
	index_a := strings.Index(equation, "x^2") // len 3
	raw_a := equation[:index_a]
	a,_ = strconv.Atoi(raw_a)

	start_index_b := index_a + 3
	index_b := strings.Index(equation[start_index_b:], "x")
	raw_b := equation[start_index_b : start_index_b+index_b]
	b, _ = strconv.Atoi(string(raw_b))
	start_index_c := start_index_b + index_b + 1
	raw_c := equation[start_index_c:strings.Index(equation, "=")]

	c, _ = strconv.Atoi(string(raw_c))
	raw_rigth_c := equation[strings.Index(equation, "=")+1:]
	right_c,_ = strconv.Atoi(string(raw_rigth_c))
	c = c-right_c
	
	// Printf("%dx^2\n",  a)
	// Printf("%dx\n", b)
	// Printf("%d\n", c)

	delta := math.Sqrt(float64(b*b -4*a*c))

	x0 := (-float64(b) + delta)/float64(2*a)
	x1 := (- float64(b) - delta)/float64(2*a)
	if delta > 0{
		Printf("Esistono due soluzioni reali: %.3f e %.3f\n",x0,x1)
	}else if delta == 0{
		Printf("Esiste un'unica soluzione reale: %.3f\n",x0)
	}else{
		Println("Non ci sono soluzioni reali")
	}

	if x0 - soglia > epsilon{
		Printf("La soluzione %.3f è maggiore della soglia",x0)
	}
	if x1 - soglia > epsilon{
		Printf("La soluzione %.3f è maggiore della soglia",x1)
	}

	_ = soglia
	_ = epsilon

}
