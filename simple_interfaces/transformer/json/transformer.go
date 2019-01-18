package json

import (
	"strconv"

	"github.com/thanosKontos/go-playground/simple_interfaces/article"
)

type Json struct{}

func New() Json {
	return Json{}
}

func (j Json) Transform(a *article.Article) string {
	return "{\"id\": " + strconv.Itoa(a.Id) + ", \"description\": \"" + a.Description + "\"}"
}
