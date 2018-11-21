package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var t int
	var s string

	stdin, err := os.Open("B-large-practice.in")

	if err != nil {
		log.Fatal(err)
	}

	os.Stdin = stdin

	fmt.Scan(&t)
	for m := 0; m < t; m++ {
		fmt.Scan(&s)
		ans := loop(s)
		fmt.Printf("case %v: %v\n", m+1, ans)
	}

}

func loop(s string) int {
	b := strings.Split(s, "")
	cur := b[0]
	ans := 0
	for i := 0; i < len(b); i++ {
		if b[i] != cur {
			ans++
			cur = b[i]
		}
	}
	if cur == "-" {
		ans++
	}
	return ans
}
