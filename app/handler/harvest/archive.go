package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nurfan/academic-literature-crawler/adapter/oaipmh"
	"github.com/nurfan/academic-literature-crawler/app/repo"
	w "github.com/nurfan/academic-literature-crawler/app/worker"
	m "github.com/nurfan/academic-literature-crawler/constants/model"
	"github.com/nurfan/academic-literature-crawler/lib/workerpool"
	"github.com/olivere/elastic/v7"
)

// HarvestArchive initiate object
type HarvestArchive struct {
	arcRepo repo.ArchiveElasticRepo
}

// Handle : handle request for this action
func (ha *HarvestArchive) Handle(c echo.Context) (err error) {

	repo := strings.ToUpper(c.Param("repo"))
	oaiClient := oaipmh.NewClient()
	oaiClient.SetRepository(repo)
	//oaiClient.SetDateRange("2020-08-01", "2020-12-31")

	url := oaiClient.Request.BaseURL
	seeders := strings.Split(url, ",")

	for i := 0; i < len(seeders); i++ {
		go ha.callSender(oaiClient, repo, seeders[i])
	}

	return c.JSON(http.StatusCreated, "Tested Create Article PASS")
}

func (ha *HarvestArchive) callSender(oai *oaipmh.OAI, platform, url string) error {
	var res *m.OaiResponse
	resumeToken := make(map[string]string)
	resumeToken[url] = ""

	x := 0
	for {

		if resumeToken[url] == "" && x > 0 {
			break
		}

		if x > 0 {
			oai.Request.MetadataPrefix = ""
		}

		log.Println("iterasi : ", x)
		oai.Request.BaseURL = url
		res, _ = oai.GetOAI()

		records := res.GetListRecord()

		resumeToken[url] = res.ListRecords.ResumptionToken.Text
		oai.Request.ResumptionToken = resumeToken[url]
		log.Println("#### Resume Token : ", resumeToken[url])

		x++

		harvesWork := &w.HarvesWork{
			Platform: platform,
			Records:  records,
			ArcRepo:  ha.arcRepo,
		}
		work := workerpool.Job{Executor: harvesWork}
		workerpool.JobQueue <- work
	}

	return nil
}

// NewHarvestArchive setup initiate object
func NewHarvestArchive(elasticConn *elastic.Client) *HarvestArchive {
	return &HarvestArchive{
		arcRepo: repo.NewArchiveIndex(elasticConn),
	}
}
