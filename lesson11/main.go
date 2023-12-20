package main

import (
	"bufio"
	. "fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//ripetizioni()
	//ppm()
	subSequence()
}

func subSequence(){

}

type RGB struct {
	r, g, b uint
}
type Img struct {
	w, h  int
	pixel []RGB
}

func ppm() {
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
			
			var img Img = Img{width,heigth,[]RGB{}}
			_ = img
			for i := 3; i < len(lines)-1; i++ {
				colors := strings.Split(lines[i]," ")
				Println(colors)
				r,_ := strconv.Atoi(colors[0])
				g,_ := strconv.Atoi(colors[1])
				b,_ := strconv.Atoi(colors[2])
				pixel := RGB{uint(r),uint(g),uint(b)}
				img.pixel = append(img.pixel,pixel)
			}
			Println(img)

		}
	

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
