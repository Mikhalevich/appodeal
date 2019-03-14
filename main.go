package main

import (
	"bufio"
	"fmt"
	"os"
)

func sortChars(line []byte) []byte {
	chars := [26]int{}
	for _, b := range line {
		chars[b-'a']++
	}

	res := make([]byte, 0, len(line))
	for i, n := range chars {
		for j := 0; j < n; j++ {
			res = append(res, byte(i)+'a')
		}
	}
	return res
}

func findAnagrams(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	res := make(map[string][]string)

	r := bufio.NewReader(file)
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			break
		}
		key := string(sortChars(l))
		anagrams := res[key]
		anagrams = append(anagrams, string(l))
		res[key] = anagrams
	}

	for _, v := range res {
		if len(v) > 1 {
			fmt.Println(v)
		}
	}

	return err
}

func main() {
	if err := findAnagrams("data.txt"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done...")
}
