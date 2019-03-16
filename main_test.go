package main

import "testing"

func BenchmarkASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := findAnagrams("small_ascii.txt", ASCIISorter{}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := findAnagrams("small_unicode.txt", UnicodeSorter{}); err != nil {
			b.Fatal(err)
		}
	}
}
