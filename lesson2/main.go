package main

import . "fmt"
import "math"

func main() {
	//rectangle()
	//circle()
	//km()
	//age()
	//area()
	//time()
}

func time() { // 3600 = 1h // 60s = 1m
	var sec int
	Print("Enter seconds: ")
	Scan(&sec)
	Print(sec)
	var h, m int
	h = sec / 3600
	sec %= 3600
	m = sec / 60
	sec %= 60

	Print("h:m:s - ", h, ":", m, ":", sec)
}
func area() {
	var n int
	var len float64
	Print("Enter n. lati: ")
	Scan(&n)
	Print("Enter length: ")
	Scan(&len)
	var area = (float64(n) * math.Pow(len, 2)) / (4 * math.Tan(math.Pi/float64(n)))
	Println("Area: ", area)
}
func age() {
	var a, b int
	Print("Enter first age: ")
	Scan(&a)
	Print("Enter second age: ")
	Scan(&b)
	mean := (float64(a+b) / 2)
	Println("Sum:", a+b, "\nMean:", mean, "\nFloor Mean:", math.Floor(mean), "\nCeil Mean:", math.Ceil(mean), "\nFuture 10 yeas: Sum:", a+b+20, "\t Mean:", (float64(a+b+20) / 2))
}
func km() {
	var km float64
	Print("distance (km): ")
	Scan(&km)
	Println(km * 0.62137)
}
func rectangle() {

	var b, h int
	Print("Enter base: ")
	Scan(&b)
	Print("Enter heigth: ")
	Scan(&h)
	Println("2p =", (b+h)*2)
	Println("A =", b*h)
}
func circle() {
	var r float64
	Print("Enter Raggio: ")
	Scan(&r)
	Println("circonferenza =", r*2*math.Pi)
	Println("A =", r*r*math.Pi)
}
