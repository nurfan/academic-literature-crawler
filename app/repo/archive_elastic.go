package repo

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	guuid "github.com/google/uuid"
	m "github.com/nurfan/academic-literature-crawler/constants/model"
	s "github.com/nurfan/academic-literature-crawler/constants/state"
	"github.com/nurfan/academic-literature-crawler/lib/errors"
	"github.com/olivere/elastic/v7"
)

// ArchiveIndex :
type ArchiveIndex struct {
	e  *errors.Error
	db *elastic.Client
}

// Create get data accounts by account number
func (c *ArchiveIndex) Create(ctx context.Context, platform string, content m.Record) (resp *elastic.IndexResponse, err error) {
	dc := content.Metadata.Dc
	dc.Identifier = append(dc.Identifier, content.Header.Identifier)
	var relation string

	uid := guuid.New().String()

	if platform == "OJS" {
		tmpStr := mergeDC(dc.Relation)
		relation = strings.ReplaceAll(tmpStr, "/view/", "/download/")
	}

	doc := m.Archive{
		ArchiveID:     uid,
		Platform:      platform,
		OaiIdentifier: content.Header.Identifier,
		Title:         mergeDC(dc.Title),
		Creator:       mergeDC(dc.Creator),
		Subject:       mergeDC(dc.Subject),
		Description:   mergeDC(dc.Description),
		Publisher:     mergeDC(dc.Publisher),
		Contributor:   mergeDC(dc.Contributor),
		Date:          mergeDC(dc.Date),
		Type:          mergeDC(dc.Type),
		Identifier:    mergeDC(dc.Identifier),
		Language:      mergeDC(dc.Language),
		Rights:        mergeDC(dc.Rights),
		Format:        mergeDC(dc.Format),
		Source:        mergeDC(dc.Source),
		Relation:      relation,
		Coverage:      mergeDC(dc.Coverage),
	}

	log.Println("Create Elastic Doc in Archive Index", uid)

	resp, err = c.db.Index().
		Index("archives").
		Id(uid).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		// Handle error
		log.Println("ERROR Create doc : ", err)
		return nil, err
	}

	return
}

// Search for search archive
func (c *ArchiveIndex) Search(ctx context.Context, platform, page, key string) (*elastic.SearchResult, int, error) {
	var currentPage, from, pageSize int = 0, 0, 10

	currentPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err)
		currentPage = 1
	}

	if currentPage > 1 {
		from = (currentPage * pageSize) - 1
	}

	//Search with a term query
	termQuery1 := elastic.NewMultiMatchQuery(key, "title", "creator", "subject", "publisher").Type("phrase_prefix")
	termQuery2 := elastic.NewMatchQuery("platform", platform)

	search := c.db.Search().Index("archives")

	if platform != "" {
		qBool := elastic.NewBoolQuery()
		esQuery := qBool.Must(termQuery1, termQuery2)

		search.Query(esQuery)
	} else {
		search.Query(termQuery1)
	}
	searchResult, err := search.From(from).Size(pageSize).Do(ctx)

	return searchResult, currentPage, nil
}

// SearchByArchiveID for search archive
func (c *ArchiveIndex) SearchByArchiveID(ctx context.Context, key string) (*elastic.SearchResult, error) {

	//Search with a term query
	termQuery := elastic.NewMultiMatchQuery(key, "archive_id").Type("phrase_prefix")
	searchResult, err := c.db.Search().
		Index("archives").
		Query(termQuery).
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

	return searchResult, nil
}

func mergeDC(param []string) (result string) {
	if len(param) > 1 {
		result = strings.Join(param[:], s.OAI_SEPARATOR)
		return
	}

	result = strings.Join(param[:], "")
	return
}

// NewArchiveIndex create new instance of ArchiveIndex
func NewArchiveIndex(db *elastic.Client) *ArchiveIndex {
	return &ArchiveIndex{
		db: db,
	}
}
