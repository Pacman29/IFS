package main

import "fmt"

func main() {
	str, find := "ABAAAABAACD", "ABA"

	fmt.Printf("NaiveStringMatcher: %s, %s, %v\n", str, find, NaiveStringMatcher(str, find))
	fmt.Printf("RabinKarpMatcher: %s, %s, %v\n", str, find, RabinKarpMatcher(str, find))
	fmt.Printf("FAMatcher: %s, %s, %v\n", str, find, FAMatcher(str, find))
	fmt.Printf("KnuthMorrisPratMatcher: %s, %s, %v\n", str, find, KnuthMorrisPratMatcher(str, find))
	fmt.Printf("BoyerMooreMatcher: %s, %s, %v\n", str, find, BoyerMooreMatcher(str, find))
}
