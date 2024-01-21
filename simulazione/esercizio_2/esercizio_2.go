package main

import (
	. "fmt"
	"os"
	"sort"
)

func main() { //1h 39min
	var input string = os.Args[1]
	string_count := Sottostringhe(input)
	sorted_count := map_to_sorted_slice(string_count)
	stampa_string_count(sorted_count, string_count)

}
func map_to_sorted_slice(string_count map[string]int) (sorted []string) {
	for k := range string_count {
		sorted = append(sorted, k)
	}
	sort.StringSlice.Sort(sorted)
	return
}

func stampa_string_count(sorted []string, string_count map[string]int) {
	Printf("output:\n")
	for i := 0; i < len(sorted); i++ {
		Printf("%s %d\n", sorted[i], string_count[sorted[i]])
	}
}

func Sottostringhe(s string) map[string]int {
	var strings_count map[string]int = map[string]int{}
	for begin := 0; begin < len(s)-1; begin++ {
		if (s[begin] < 97 || s[begin] > 122) && (s[begin+1] < 97 || s[begin+1] > 122){
			return map[string]int{}
		}
		for end := begin + 1; end < len(s); end++ {
			if s[begin] < s[end] {
				strings_count[s[begin:end+1]]++
			} else {
				break
			}
		}
	}
	return strings_count
}
