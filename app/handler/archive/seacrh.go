package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nurfan/academic-literature-crawler/app/repo"
	m "github.com/nurfan/academic-literature-crawler/constants/model"
	"github.com/olivere/elastic/v7"
)

// SearchArchive initiate object
type SearchArchive struct {
	elastic *elastic.Client
	arcRepo repo.ArchiveElasticRepo
}

// Handle : handle request for this action
func (ha *SearchArchive) Handle(c echo.Context) (err error) {
	var result m.APIResponse

	ctx := c.Request().Context()
	key := c.QueryParam("keyword")
	page := c.QueryParam("page")
	platform := c.QueryParam("platform")

	searchResult, currentPage, err := ha.arcRepo.Search(ctx, platform, page, key)

	if err == nil {
		var result m.APIResponse
		var listArchive []m.ListArchive

		if searchResult.TotalHits() > 0 {

			for _, hit := range searchResult.Hits.Hits {
				var t m.ListArchive
				err := json.Unmarshal(hit.Source, &t)
				if err == nil {
					log.Println(err)
				}
				t.Link = os.Getenv("HOST") + "/api/v1/archive/" + strings.ToLower(t.Platform) + "/" + t.ArchiveID
				listArchive = append(listArchive, t)
			}

			result.Code = http.StatusOK
			result.Message = http.StatusText(result.Code)
			result.Data = map[string]interface{}{
				"response_time": searchResult.TookInMillis,
				"current_page":  currentPage,
				"total_result":  searchResult.TotalHits(),
				"archive":       listArchive,
			}

			return c.JSON(result.Code, result)
		}
	}

	result.Code = http.StatusNotFound
	result.Message = http.StatusText(result.Code)

	return c.JSON(result.Code, result)
}

// NewSearchArchive setup initiate object
func NewSearchArchive(elasticConn *elastic.Client) *SearchArchive {
	return &SearchArchive{
		elastic: elasticConn,
		arcRepo: repo.NewArchiveIndex(elasticConn),
	}
}
