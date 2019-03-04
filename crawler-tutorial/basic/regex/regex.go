package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is guopuke@gmail.com@aaaa.com
emaiasd1 is abbbb@a..com
emaiasd2 is 22afsfasb@B.com
emaiasd2 is 1111@B.com.org
`

func main() {
	compile := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := compile.FindAllString(text, -1)
	submatch := compile.FindAllStringSubmatch(text, -1)

	fmt.Println("Find All String Submatch\n")
	for _, m := range submatch {
		fmt.Println(m)
	}

	fmt.Println("\n\nFind All String\n")
	fmt.Println(match)
}
