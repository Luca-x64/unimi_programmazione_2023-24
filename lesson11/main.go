package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// ripetizioni()
	// ppm()
	// subSequence()
	pitagorica()
}

func pitagorica(){
	n,_ := strconv.Atoi(os.Args[1])
	tavola := CreaTavolaPitagorica(n)
	StampaTavolaPitagorica(tavola)
}

func CreaTavolaPitagorica(n int) [][]int{
	var tavola [][]int
	for i := 1; i <= n; i++ {
		var line []int
		for j := 1; j <= n; j++ {
			line = append(line,i*j)
		}
		tavola = append(tavola,line)
	}
	return tavola
}

func StampaTavolaPitagorica(s [][]int){
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			Printf("%4d ",s[i][j])
		}
		Println()
	}

}

func subSequence(){
	seq := os.Args[1:]
	var subsequences []string 
	for s := 0;s  < len(seq)/2; s++ {
		for e := s+1; e < len(seq); e++ {
			if seq[s] == seq[e]{
				subsequences = append(subsequences, strings.Join((seq[s:e+1])," ") )
				
			}
		}
	}
	slices.SortFunc(subsequences,func(a,b string) int {
		return cmp.Compare(len(a),len(b))
	})
	for i := 0; i < len(subsequences); i++ {
		Println(subsequences[i])
	}
}

type RGB struct {
	r, g, b uint
}
type Img struct {
	w, h  int
	pixels []RGB
}

func ppm() {
	img := LeggiFilePPM()
	Printf("Immagine letta: %v",img)
	
}

func LeggiFilePPM() (img Img){
	file, err := os.Open(os.Args[1])
	if err == nil {
		b := make([]byte, 4096)
			n, err := file.Read(b)
			if err != nil {
				file.Close()
				//break
			}
			data := string(b[:n])
	
			lines := strings.Split(data,"\n")
			_ = lines[0]
			_ = lines[2]
			sizes := strings.Split(lines[1],"")
			width,_ := strconv.Atoi(sizes[0])
			heigth,_ := strconv.Atoi(sizes[1])
			img.w=width
			img.h = heigth
			for i := 3; i < len(lines)-1; i++ {
				colors := strings.Split(lines[i]," ")
				// Println(colors)
				r,_ := strconv.Atoi(colors[0])
				g,_ := strconv.Atoi(colors[1])
				b,_ := strconv.Atoi(colors[2])
				pixel := RGB{uint(r),uint(g),uint(b)}
				img.pixels = append(img.pixels,pixel)
			}
		}
		return

}

func ripetizioni() {
	scan := bufio.NewScanner(os.Stdin)
	mappa := map[string]int{}
	_ = mappa
	for scan.Scan() {
		text := scan.Text()
		split := strings.Split(text, " ")
		for s := range split {
			word := strings.ToUpper(split[s])
			if len(word) > 0 {
				var valid bool = true
				for c := range word {
					if unicode.IsSpace(rune(word[c])) {
						valid = false
					}
				}

				if valid {
					mappa[word] += 1
				}
			}
		}

	}
	list := []string{}
	for k := range mappa {
		list = append(list, k)
	}
	slices.Sort(list)
	for i := range list {
		Printf("%s: %d\n", list[i], mappa[list[i]])
	}
}
