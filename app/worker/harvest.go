package worker

import (
	"context"
	"log"

	"github.com/nurfan/academic-literature-crawler/app/repo"
	m "github.com/nurfan/academic-literature-crawler/constants/model"
)

// HarvesWork job data
type HarvesWork struct {
	Platform string
	Records  []m.Record
	ArcRepo  repo.ArchiveElasticRepo
}

//Handle THIS IS WHERE PROCESSING WILL RUNNING, EVERY STRUCT MUST HAVE FUNCTION HANDLE
func (h *HarvesWork) Handle() error {

	ctx := context.Background()
	for _, record := range h.Records {
		if record.Header.Status != "deleted" {
			_, _ = h.ArcRepo.Create(ctx, h.Platform, record)
		} else {
			log.Println("Failed create document, article was deleted")
		}
	}
	return nil
}
