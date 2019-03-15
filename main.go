package main

import (
	"bufio"
	"fmt"
	"os"
)

// sorting characters in bytearray and returns a sorted copy with all letters mapped to their lower case
// this function make deal with a-zA-z characters set, any other character will trigger error
func sortCharsASCII(b []byte) ([]byte, error) {
	chars := [26]int{}
	for _, c := range b {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}

		if c < 'a' || c > 'z' {
			return nil, fmt.Errorf("Invalid character: %c in %s", c, string(b))
		}
		chars[c-'a']++
	}

	res := make([]byte, 0, len(b))
	for i, n := range chars {
		for j := 0; j < n; j++ {
			res = append(res, byte(i)+'a')
		}
	}
	return res, nil
}

// parse file and returns slice of agagrams
func findAnagrams(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	anagramCache := make(map[string][]string)

	r := bufio.NewReader(file)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			break
		}

		key, err := sortCharsASCII(l)
		if err != nil {
			return nil, err
		}
		anagrams := anagramCache[string(key)]
		anagrams = append(anagrams, string(l))
		anagramCache[string(key)] = anagrams
	}

	var res [][]string
	for _, v := range anagramCache {
		if len(v) > 1 {
			res = append(res, v)
		}
	}

	return res, err
}

func main() {
	anagrams, err := findAnagrams("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(anagrams)
	fmt.Println("Done...")
}
