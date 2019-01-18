package simpletext

import "github.com/thanosKontos/go-playground/simple_interfaces/article"

type Simpletext struct{}

func New() Simpletext {
	return Simpletext{}
}

func (s Simpletext) Transform(a *article.Article) string {
	return a.Description
}
