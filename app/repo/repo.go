package repo

import (
	"context"

	m "github.com/nurfan/academic-literature-crawler/constants/model"
	"github.com/olivere/elastic/v7"
)

// ArchiveElasticRepo : archive index abstract
type ArchiveElasticRepo interface {
	Create(context.Context, string, m.Record) (*elastic.IndexResponse, error)
	Search(context.Context, string, string, string) (*elastic.SearchResult, int, error)
	SearchByArchiveID(context.Context, string) (*elastic.SearchResult, error)
}
