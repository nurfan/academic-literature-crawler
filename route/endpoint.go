package route

import (
	article "github.com/nurfan/academic-literature-crawler/article/handler"
	harvest "github.com/nurfan/academic-literature-crawler/harvest/handler"

	"github.com/labstack/echo/v4"
)

// Handler endpoint to use it later
type Handler interface {
	Handle(c echo.Context) (err error)
}

var endpoint = map[string]Handler{

	// article
	"create_article": article.NewCreateArticle(),
	"read_article":   article.NewReadArticle(),

	//harvest
	"harvest_archive": harvest.NewHarvestArchive(),
}
