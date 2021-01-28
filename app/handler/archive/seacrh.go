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
	var pageNUmber, pageSize int = 0, 15

	// create context
	ctx := c.Request().Context()

	key := c.QueryParam("keyword")
	page := c.QueryParam("page")

	pageNUmber, err = strconv.Atoi(page)
	if err != nil {
		pageNUmber = 0
	}

	pageNUmber = (pageNUmber + 1) * pageSize

	// Search with a term query
	termQuery := elastic.NewMultiMatchQuery(key, "title", "creator", "subject", "description", "publisher", "source").Type("phrase_prefix")
	searchResult, err := ha.elastic.Search().
		Index("archives").
		Query(termQuery).
		From(pageNUmber).Size(15).
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

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	//fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.

	// TotalHits is another convenience function that works even when something goes wrong.
	//fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	//var ttyp m.Archive
	// /searchResult.Each(reflect.TypeOf(ttyp))

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

		result.Code = http.StatusFound
		result.Message = http.StatusText(result.Code)
		result.Data = map[string]interface{}{
			"response_time": searchResult.TookInMillis,
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
