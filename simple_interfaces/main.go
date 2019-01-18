package main

import (
	"fmt"

	"github.com/thanosKontos/go-playground/simple_interfaces/article"
	"github.com/thanosKontos/go-playground/simple_interfaces/transformer/json"
)

func main() {
	j := json.New()
	a := &article.Article{3, "abcd"}

	fmt.Print(j.Transform(a))
}
