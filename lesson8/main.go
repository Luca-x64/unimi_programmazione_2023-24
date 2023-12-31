package main

import (
	. "fmt"
	_ "math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	_ "strings"
	"time"
)

func main() {
	//SumOfEven()
	//filterVoti()
	//randomNumbers()
	//sumUnique()
	//sumOfProduct()
	//Prime()
}

func Prime() {
	n, err := strconv.Atoi(os.Args[1])
	_ = n
	sn := os.Args[1]
	var nums []int
	if isprime(n) {
		nums = append(nums, n)
	}
	if err == nil {
		for i := 0; i <= 3; i++ { // 1
			for j := len(sn); j > i-1; j-- {
				a, _ := strconv.Atoi(sn[:j-i] + sn[j:])
				if isprime(a) { // filter prime
					nums = append(nums, a)
				}
			}
		}
		sort.Ints(nums)
		for _, v := range nums {
			Println(v)
		}
	} else {
		Print("args error")
	}
}

func isprime(n int) (prime bool) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumOfProduct() {
	nums := leggiNumeri()
	sum := 0
	for i := 0; i < len(nums)-1; i += 2 {
		sum += nums[i] * nums[i+1]
	}
	Printf("Sum: %d", sum)
}

func sumUnique() {
	nums := leggiNumeri()
	sum := 0
	for _, v := range nums {
		if occorrenze(nums, v) == 1 {
			sum += v
		}
	}
	Printf("sum: %d", sum)
}
func occorrenze(numeri []int, n int) (cnt int) {
	for _, v := range numeri {
		if v == n {
			cnt++
		}
	}
	return
}

func randomNumbers() {
	soglia, _ := strconv.Atoi(os.Args[1])
	gen := genera(soglia)
	filtered := gen[:len(gen)-1]
	Println("generati:", gen, "\nValori sopra la soglia:", filtered)
}

func genera(s int) []int {
	var rnd int = 100
	var gen []int
	for rnd >= s {
		rand.Seed(int64(time.Now().Nanosecond()))
		rnd = rand.Intn(102) - 1
		gen = append(gen, rnd)
	}
	return gen
}

func filterVoti() {
	Println(filtraVoti(leggiNumeri()))
}

func filtraVoti(voti []int) (sufficienti, insufficienti []int) {
	for _, v := range voti {
		if v < 60 {
			insufficienti = append(insufficienti, v)
		} else {
			sufficienti = append(sufficienti, v)
		}
	}
	return
}

func leggiNumeri() (numeri []int) {
	args := os.Args[1:]
	for _, v := range args { //parse int args
		n, _ := strconv.Atoi(v)
		numeri = append(numeri, n)
	}
	return
}

func SumOfEven() {
	args := os.Args[1:]
	var nums []int
	for _, v := range args { //parse int args
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	Println(calcolaEven(nums))
}

func calcolaEven(sl []int) (sum int) {
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			product := sl[i] * sl[j]
			if product%2 == 0 {
				sum += product
			}
		}
	}
	return
}
