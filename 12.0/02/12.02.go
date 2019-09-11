package main

import (
	"fmt"

	"github.com/Golang/12.0/02/quote"
	"github.com/Golang/12.0/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
