package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	var currentPage, pageSize, from int = 0, 15, 0

	// create context
	ctx := c.Request().Context()

	key := c.QueryParam("keyword")
	page := c.QueryParam("page")

	currentPage, err = strconv.Atoi(page)
	if err != nil {
		log.Println(err)
		currentPage = 1
	}

	log.Println(currentPage)

	if currentPage >= 1 {
		from = (currentPage * pageSize)
	}

	// Search with a term query
	termQuery := elastic.NewMultiMatchQuery(key, "title", "creator", "subject", "publisher", "source", "platform").Type("phrase_prefix")
	searchResult, err := ha.elastic.Search().
		Index("archives").
		Query(termQuery).
		From(from).Size(15).
		Pretty(true).
		Do(ctx)

	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			panic(fmt.Sprintf("Document not found: %v", err))
		case elastic.IsTimeout(err):
			panic(fmt.Sprintf("Timeout retrieving document: %v", err))
		case elastic.IsConnErr(err):
			panic(fmt.Sprintf("Connection problem: %v", err))
		default:
			// Some other kind of error
			panic(err)
		}
	}

	var result m.APIResponse

	var t m.ListArchive
	var listArchive []m.ListArchive
	// Here's how you iterate through results with full control over each step.
	if searchResult.TotalHits() > 0 {

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).

			err := json.Unmarshal(hit.Source, &t)
			if err == nil {
				log.Println(err)
			}

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
