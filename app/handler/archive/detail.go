package handler

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"reflect"

// 	"github.com/labstack/echo/v4"
// 	"github.com/nurfan/academic-literature-crawler/app/repo"
// 	m "github.com/nurfan/academic-literature-crawler/constants/model"
// 	"github.com/olivere/elastic/v7"
// )

// // DetailArchive initiate object
// type DetailArchive struct {
// 	elastic *elastic.Client
// 	arcRepo repo.ArchiveElasticRepo
// }

// // Handle : handle request for this action
// func (da *DetailArchive) Handle(c echo.Context) (err error) {
// 	ctx := c.Request().Context()

// 	idDoc := c.Param("idDoc")

// 	// Search with a term query
// 	//termQuery := elastic.NewMultiMatchQuery("Soil", "title", "creator", "subject", "description", "publisher", "source").Type("phrase_prefix")
// 	searchResult, err := da.elastic.Search().
// 		Index("archives").
// 		Type("_doc").
// 		Pretty(true).
// 		Do(context.Background())

// 	if err != nil {
// 		// Handle error
// 		log.Println(err)
// 	}

// 	// searchResult is of type SearchResult and returns hits, suggestions,
// 	// and all kinds of other information from Elasticsearch.
// 	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

// 	// Each is a convenience function that iterates over hits in a search result.
// 	// It makes sure you don't need to check for nil values in the response.
// 	// However, it ignores errors in serialization. If you want full control
// 	// over iterating the hits, see below.
// 	var ttyp m.Archive
// 	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
// 		t := item.(m.Archive)
// 		log.Println(t)
// 	}
// 	// TotalHits is another convenience function that works even when something goes wrong.
// 	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

// 	return c.JSON(http.StatusCreated, "Tested Create Article PASS")
// }

// //NewDetailArchive setup initiate object
// func NewDetailArchive(elasticConn *elastic.Client) *DetailArchive {
// 	return &DetailArchive{
// 		elastic: elasticConn,
// 		arcRepo: repo.NewArchiveIndex(elasticConn),
// 	}
// }
