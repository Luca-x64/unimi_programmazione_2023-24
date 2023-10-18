package main

import (
	. "fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//segno()
	//multiplo10()
	//intervallo()
	//fizzBuzz()
	//even()
	//division()
	//angoli()
	//conversions()
	//retta()
	//tabellina()
	//sumInterval()
	//operationN()
	//mean() // TODO check scanf input stop condition
	//guess()
	//fizz2Buzz()
	//divisors()
	//operations()
	//crescente() // TODO check scanf 
	//fibonacci()
}

func fibonacci() {
	var a, b int = 1, 1
	Println("enter number: ")
	Scan(&n)
	Print(a, " ")
	for i := 1; i < n; i++ {
		Print(b, " ")
		a, b = b, b+a
	}
}

func crescente() {
	Println("Inserisci una sequenza di numeri: ")
	for false { // debug

	}
}

func operations() {
	var a, b int
	Print("Enter 2 numbers: ")
	Scan(&a, &b)
	if a > b {
		Println("Maggiore: ", a, "\nMinore: ", b)
	} else {
		Println("Maggiore: ", b, "\nMinore: ", a)
	}
	Println("Somma: ", a+b, "\nDifferenza: ", math.Abs(float64(a-b)), "\nProdotto: ", a*b, "\nDivisione: ", float64(a)/float64(b), "\nMean: ", float64((a+b))/2)
	var potenzaFor int = 1
	for i := 1; i <= b; i++ {
		potenzaFor *= a
	}
	Println("Potenza for: ", potenzaFor, "\nPotenza math: ", math.Pow(float64(a), float64(b)))

}

func divisors() {
	var n int
	Print("Enter number: ")
	Scan(&n)

	Printf("Divisori di %d: ", n)
	for i := 1; i < n; i++ {
		if n%i == 0 {
			Print(i, " ")
		}
	}

}

func fizz2Buzz() {
	var n int
	Print("Enter number: ")
	Scan(&n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			Print("FizzBuzz ")
		} else if i%3 == 0 {
			Print("Fizz ")
		} else if i%5 == 0 {
			Print("Buzz ")
		} else {
			Print(i, " ")
		}
	}
}

func guess() {
	rand.Seed(int64(time.Now().Nanosecond()))
	var rnd int = rand.Intn(100)
	var input int
	cnt := 0
	for rnd != input {
		Print("Guess number: ")
		Scan(&input)
		if input < rnd {
			Println("Too low")
		} else if input > rnd {
			Println("Too high")
		}
		cnt++
	}
	Printf("You guess in %d attempts\n", cnt)

}

func mean() {
	var sum, cnt int
	Print("Enter number: ")
	for {
		var t int
		Scanf("%d", &t) // dont work on arch linux
		if t > 0 {
			sum += t
			cnt++
		} else {
			break
		}
	}
	Printf("Mean=%f \n", float64(sum)/float64(cnt))
}

func operationN() {
	var n int
	Print("Enter number: ")
	Scan(&n)

	max, min, gt, eq, ls, sum := 0, 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		var t int
		Print("Enter number: ")
		Scan(&t)
		if t > max {
			max = t
		} else if t < min {
			min = t
		}

		if t > 0 {
			gt++
		} else if t == 0 {
			eq++
		} else {
			ls++
		}
		sum += t
	}
	Printf("Somma = %d\nmax = %d\nmin = %d\ninteri > 0 = %d\n,interi < 0 =%d\n,interi = 0 =%d", sum, max, min, gt, ls, eq)
}

func sumInterval() {
	var a, b, sum int
	Print("Enter 2 numbers: ")
	Scan(&a, &b)
	if a < b {
		if a%2 != 0 {
			a += 1
		}
		if b%2 != 0 {
			b -= 1
		}
		for i := a; i <= b; i += 2 {
			sum += i
		}
		Println("Somma:", sum)
	}
}

func tabellina() {
	var n int
	Print("Enter number: ")
	Scan(&n)
	for i := 1; i <= 10; i++ {
		Println(i, "x", n, "=", i*n)
	}
}

func retta() {
	var m, q, px, py float64
	Print("Enter m and q:")
	Scan(&m, &q)
	Print("Enter x and y:")
	Scan(&px, &py)
	y := m*px + q
	if y < py {
		Println("Il punto è sopra alla retta")
	} else if y == py {
		Println("Il punto appartiene alla retta")
	} else {
		Println("Il punto è sotto alla retta")
	}
}

func conversions() {
	var choise, value int
	Print("Scegli la conversione:\n1: secondi (inseriti dall’utente) in ore n 2: secondi inseriti dall’utente in minuti \n3: minuti inseriti dall’utente in ore \n4: minuti inseriti dall’utente in secondi \n5: ore inserite dall’utente in secondi \n6: ore inserite dall’utente in minuti \n7: minuti inseriti dall’utente in giorni e ore \n8: minuti inseriti dall’utente in anni e giorni\n:")
	Scan(&choise)
	Print("Inserisci il valore da convertire: ")
	Scan(&value)
	switch choise {
	case 1:
		Println(value, "secondi corrispondono a ", float64(value)/3600, "ore")
	case 2:
		Println(value, "secondi corrispondono a ", float64(value)/60, "minuti")
	case 3:
		Println(value, "minuti corrispondono a ", float64(value)/60, "ore")
	case 4:
		Println(value, "minuti corrispondono a ", value*60, "secondi")
	case 5:
		Println(value, "ore corrispondono a ", value*3600, "secondi")
	case 6:
		Println(value, "ore corrispondono a ", value*60, "minuti")
	case 7:
		giorni := value / 1440
		Println(value, "minuti corrispondono a ", giorni, "giorni e ", float64((value%(giorni*1440)))/60, " ore")
	case 8:
		anni := value / 525600
		Println(value, "minuti corrispondono a ", anni, "giorni e ", float64((value%(anni*525600)))/1440, " giorni")
	default:
		Println("Error choise")
	}
}

func angoli() {
	var n1, n2 int
	Print("Enter 2 angles of a trinalge (degree): ")
	Scan(&n1, &n2)
	n3 := 180 - n1 - n2
	if n3 > 0 {
		Println("3rd angle:", n3)
	} else {
		Println("I due angoli non appartengono ad un triangolo")
	}

}

func division() {
	var n1, n2 int
	Print("Enter 2 numbers: ")
	Scan(&n1, &n2)
	if n2 > 0 {
		Println("Quoziente: ", (float64(n1) / float64(n2)))
	} else {
		Println("Impossibile")
	}
}

func even() {
	var n int
	Print("Enter number: ")
	Scan(&n)
	if n%2 == 0 {
		Println("Even")
	} else {
		Println("Odd")
	}
}

func fizzBuzz() {
	var n int
	Print("Enter number: ")
	Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		Println("Fizz Buzz")
	} else if n%3 == 0 {
		Println("Fizz")
	} else if n%5 == 0 {
		Println("Buzz")
	} else {
		Println()
	}
}

func intervallo() {
	var n int
	Print("Enter mark: ")
	Scan(&n)
	if n < 60 {
		Println("Insufficiente")
	} else if n >= 60 && n < 70 {
		Println("Sufficiente")
	} else if n >= 70 && n < 80 {
		Println("Buono")
	} else if n >= 80 && n < 90 {
		Println("Distinto")
	} else if n >= 90 && n < 100 {
		Println("Ottimo")
	} else {
		Println("Errore")
	}
}

func multiplo10() {
	var n int
	Print("Enter number: ")
	Scan(&n)
	if n%10 == 0 {
		Println(n, "è multiplo 10")
	} else {
		Println(n, "non è multiplo 10")
	}
}

func segno() {
	var n int
	Print("Enter number: ")
	Scan(&n)
	if n > 0 {
		Printf("+%d\n", n)
	} else {
		Println(n)
	}
}
