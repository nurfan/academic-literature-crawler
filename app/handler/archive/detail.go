package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nurfan/academic-literature-crawler/adapter/slims"
	"github.com/nurfan/academic-literature-crawler/app/repo"
	m "github.com/nurfan/academic-literature-crawler/constants/model"
	s "github.com/nurfan/academic-literature-crawler/constants/state"
	"github.com/olivere/elastic/v7"
)

// DetailArchive initiate object
type DetailArchive struct {
	elastic *elastic.Client
	arcRepo repo.ArchiveElasticRepo
}

// Handle : handle request for this action
func (ha *DetailArchive) Handle(c echo.Context) (err error) {
	var result m.APIResponse
	ctx := c.Request().Context()

	archiveID := c.Param("ID")
	searchResult, err := ha.arcRepo.SearchByArchiveID(ctx, archiveID)

	var archive m.Archive
	if searchResult.TotalHits() > 0 {

		for _, hit := range searchResult.Hits.Hits {
			err := json.Unmarshal(hit.Source, &archive)
			if err == nil {
				log.Println(err)
			}
		}

		result.Data = map[string]interface{}{
			"archive": archive,
		}

		if archive.Platform == s.SLIMS {
			bookInfo := ha.getBookInfo(archive)
			result.Data = map[string]interface{}{
				"archive": ha.mappingSlimsResponse(bookInfo),
			}
		}

		result.Code = http.StatusOK
		result.Message = http.StatusText(result.Code)

		return c.JSON(result.Code, result)
	}

	result.Code = http.StatusNotFound
	result.Message = http.StatusText(result.Code)

	return c.JSON(result.Code, result)
}

func (ha *DetailArchive) getBookInfo(book m.Archive) *m.SlimsDetailBookResponse {
	slimsID := ha.getSlimsID(book.OaiIdentifier)

	slimsClient := slims.NewClient()
	info, err := slimsClient.GetBookInfo(slimsID)

	if err != nil {
		return nil
	}

	return info
}

func (ha *DetailArchive) mappingSlimsResponse(book *m.SlimsDetailBookResponse) m.SlimsBookInformation {
	var result m.SlimsBookInformation

	author := m.Authority{
		Name: book.Mods.Name.NamePart,
		Type: book.Mods.Name.Type,
		Role: book.Mods.Name.Role.RoleTerm.Text,
	}

	result.Title = book.Mods.TitleInfo.Title
	result.Cover = os.Getenv("SLIMS_HOST") + os.Getenv("SLIMS_PATH_IMG") + book.Mods.Image
	result.Author = author
	result.PublishDate = book.Mods.OriginInfo.DateIssued
	result.Publisher = book.Mods.OriginInfo.Publisher
	result.Edition = book.Mods.OriginInfo.Edition
	result.PhysicalDescription = book.Mods.PhysicalDescription.Extent
	result.Subject = book.Mods.Subject.Topic
	result.Classification = book.Mods.Classification

	var copys []m.CopyInformation
	for _, v := range book.Mods.Location.HoldingSimple.CopyInformation {
		copy := m.CopyInformation{}

		copy.Numeration = v.NumerationAndChronology.Text
		copy.ShelfLocator = v.ShelfLocator
		copy.Sublocation = v.Sublocation

		copys = append(copys, copy)
	}

	location := m.Location{
		PhysicalLocation: book.Mods.Location.PhysicalLocation,
		ShelfLocator:     book.Mods.Location.ShelfLocator,
		CopyInformations: copys,
	}

	result.Locations = location

	return result

}

func (ha *DetailArchive) getSlimsID(str string) (wordsAfterDash string) {
	indexAfterDash := strings.Index(str, "-")

	if indexAfterDash >= 0 {
		runes := []rune(str)
		wordsAfterDash = string(runes[indexAfterDash:len(str)])
		wordsAfterDash = strings.Replace(wordsAfterDash, "-", "", -1)
	}

	return
}

// NewDetailArchive setup initiate object
func NewDetailArchive(elasticConn *elastic.Client) *DetailArchive {
	return &DetailArchive{
		elastic: elasticConn,
		arcRepo: repo.NewArchiveIndex(elasticConn),
	}
}
