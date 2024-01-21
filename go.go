package main

import (
	. "fmt"
	"math"
)

func main(){
	first()
	second()
	third()
}

func third() {
	Println(3,"--")
	if a := 0; a == 0 { // declaration; condition
		Println(a)
	}
	//Print(a) //Error, a not exist
	// for { //infinite loop  <==> for ;; //infinite
		x := 10
	for x >= 0 { // while loop
		Print(x)
		x--
	}
	Println()
	for i := 0; i < 10; i++ { // == for i:=0;i<10;{
		Print(i)
	}
	Println()
}

func second() {
	Println(2,"--")
	var a,b,c float64
	Print("Enter a b c: ")
	//Scan(&a,&b,&c)
	a,b,c  = 3,-7,4
	delta := math.Pow(b,2) -4*a*c
	x1,x2 := -b + math.Sqrt(delta)/(2*a), -b - math.Sqrt(delta/(2*a))
	Println(x1,x2,max(x1,x2))
	var f float64
	f = 1.23444
	Println("f:", f)
	f *= 1e4
	Println("f:", f)
	f1 := 1.3e+32
	if f > f1{
		Println("max:",f)
	}else{
		Println("max:",f1)
	}
	Println("f:", f)
}

func first() {
	Println(1,"--")
	Println(0b11111111) // binary, prefix = `0b`
	Println(0133,010) // oct, prefix = 0
	Println(0xc1a0) // hex, prefix `0x`
	t0 := 0
	var t1 int = 1
	var (
		t2 int = 2
		n = 0
	)
	Println(t0,t1,t2)
	//_,_,_ = t0,t1,t2
	n = n +1
	n +=1
	n++
	Println(n,n/2,n%2,n<<1,n&3,n^5,n&0b110,n^0b10,n^0x12,n^023)
	Println(-8 / -2)
}