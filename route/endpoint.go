package route

import (
	a "github.com/nurfan/academic-literature-crawler/app/handler/archive"
	h "github.com/nurfan/academic-literature-crawler/app/handler/harvest"
	"github.com/olivere/elastic/v7"

	"github.com/labstack/echo/v4"
)

// Handler endpoint to use it later
type Handler interface {
	Handle(c echo.Context) (err error)
}

var endpoint = map[string]Handler{}

func bindingConn(conn *elastic.Client) {
	endpoint = map[string]Handler{

		//harvest
		"harvest_archive": h.NewHarvestArchive(conn),

		//archive
		"search_archive": a.NewSearchArchive(conn),
		"detail_archive": a.NewDetailArchive(conn),
	}
}
