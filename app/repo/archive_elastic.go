package repo

import (
	"context"
	"log"
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

	uid := guuid.New().String()

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
		Relation:      mergeDC(dc.Relation),
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
