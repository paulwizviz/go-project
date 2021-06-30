package main

import (
	"fmt"

	"myapp/internal/mathutil"

	"github.com/paulwizviz/go-workspace/mylib"
)

func main() {

	sentence := mylib.FormSentence("Hello", "World", "and", "Earth")
	fmt.Println(sentence)

	result := mathutil.Add(1, 1)
	fmt.Println(result)
}
