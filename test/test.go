package main

import . "fmt"
import . "math"

func main()  {
	//cast32()
	//rombo()
	//test4()
	test5()
	
	
}

const Dimensione = 10

func test5() {var a []int

	for i := 0; i < Dimensione; i++ {
		a = append([]int{i + 1}, a...)
	}

	Println(a)

	b := make([]int, Dimensione)

	copy(b, a[Dimensione/2:])

	Println(b)
}

func stampa(sl []int) {
	for _, v := range sl {
		Print(v, " ")
	}
	Println()
}
func modifica(sl []int) {
	for i := range sl {
		sl[i] *= 2
	}
}

func eliminaUltimoElemento(sl []int) {
	sl = sl[:len(sl)-1]
}






func test4() {
	var n int



	Scan(&n)

	for i:=1; i<=n; i++ {

		switch {
		case i<5:
			Println(i, "Minore di 5")
			//fallthrough
		case i<10:
			Println(i, "Minore di 10")
		}

	}
}

func rombo() {
	var n int
	Print("n: ")
	Scan(&n) // 3
	//top
	for i := 0; i < n; i++ {
		for j := 0; j < n+i; j++ {
			if j < n-i-1{
				Print(" ")
			}else{
				Print("*")
			}
		}
		Println()
	}
	//bottom
	for i := n-1; i >0; i-- {
		for j := n-1 + i; j > 0; j-- {
			if j < n-i || j >= n  {
				Print(" ")
			}else{
				Print("*")
			}
		}
		Println()
	}
	
}

func cast32() {
	var a,b float64
	Scan(&a)
	for {
	  Scan(&b)
	  if b==0{
		break
	  }
	  Println(Round((a-b)*Pow(10,2)/Pow(10,2)))
	  a=b
	}
}


var a int = 10

func test1() int {
	a += 5
	return a
}

func test2(a int) int {
	a += 6
	return a
}

func test3() int {
	return a + 7
}
