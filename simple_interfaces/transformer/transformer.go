package transformer

import "github.com/thanosKontos/go-playground/simple_interfaces/article"

type Transformer interface {
	Transform(*article.Article) string
}
