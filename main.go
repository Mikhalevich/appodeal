package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type sorter interface {
	Sort(s string) (string, error)
}

// implements sorter interface for ascii characters
type ASCIISorter struct {
	// pass
}

// sorting characters in string and returns a sorted copy with all letters mapped to their lower case
// this class make deal with a-zA-z characters set, any other character will trigger error
func (as ASCIISorter) Sort(s string) (string, error) {
	chars := [26]int{}
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}

		if c < 'a' || c > 'z' {
			return "", fmt.Errorf("Invalid character: %c in %s", c, string(s))
		}
		chars[c-'a']++
	}

	res := make([]byte, 0, len(s))
	for i, n := range chars {
		for j := 0; j < n; j++ {
			res = append(res, byte(i)+'a')
		}
	}
	return string(res), nil
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// implements sorter interface for unicode characters
type UnicodeSorter struct {
	// pass
}

// sorting characters in string and returns a sorted copy with all letters mapped to their lower case
// this class make deal with unicode characters
func (us UnicodeSorter) Sort(s string) (string, error) {
	s = strings.ToLower(s)
	runes := []rune(s)

	sort.Sort(sortRunes(runes))
	return string(runes), nil
}

// parse file and returns slice of agagrams
func findAnagrams(fileName string, s sorter) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	anagramCache := make(map[string][]string)

	r := bufio.NewScanner(file)
	for r.Scan() {
		l := r.Bytes()

		key, err := s.Sort(string(l))
		if err != nil {
			return nil, err
		}
		anagrams := anagramCache[key]
		anagrams = append(anagrams, string(l))
		anagramCache[key] = anagrams
	}

	var res [][]string
	for _, v := range anagramCache {
		if len(v) > 1 {
			res = append(res, v)
		}
	}

	return res, r.Err()
}

func main() {
	anagrams, err := findAnagrams("small_ascii.txt", ASCIISorter{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ASCII:")
	fmt.Println(anagrams)

	anagrams, err = findAnagrams("small_unicode.txt", UnicodeSorter{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("UNICODE:")
	fmt.Println(anagrams)
	fmt.Println("Done...")
}
