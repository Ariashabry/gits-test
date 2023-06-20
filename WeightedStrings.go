package main

import (
	"fmt"
)

//menghitung bobot pada setiap substring
func hitungBobot(s string) []int {
	bobot := make([]int, len(s))
	count := 1
	for i := 0; i < len(s); i++ {
		bobot[i] = (int(s[i]) - 'a' + 1) * count
		if i < len(s)-1 && s[i] == s[i+1] {
			count++
		} else {
			count = 1
		}
	}
	return bobot
}

//cek apakah angka di queries bernilai sama dengan bobot karakter
func cekQueries(s string, queries []int) []string {
	bobots := hitungBobot(s)
	results := make([]string, len(queries))
	for i, query := range queries {
		found := false
		for _, bobot := range bobots {
			if bobot == query {
				results[i] = "Yes"
				found = true
				//jika queries = bobot, maka simpan sebagai true/yes dan
				//lanjutkan ke queries di index selanjutnya
				break
			}
		}
		if !found {
			results[i] = "No"
		}
	}
	return results
}

func main() {
	s := "abbcccd"
	queries := []int{1, 3, 9, 8}
	results := cekQueries(s, queries)
	fmt.Println(results)
}
