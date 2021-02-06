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
			"archive": ha.mappingOJSResponse(archive),
		}

		if archive.Platform == s.SLIMS {
			bookInfo := ha.getBookInfo(archive)
			result.Data = map[string]interface{}{
				"archive": ha.mappingSlimsResponse(archive, bookInfo),
			}
		}

		if archive.Platform == s.EPRINTS {
			result.Data = map[string]interface{}{
				"archive": ha.mappingEprintsResponse(archive),
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

func (ha *DetailArchive) mappingSlimsResponse(archive m.Archive, book *m.SlimsDetailBookResponse) m.SlimsBookInformation {
	var result m.SlimsBookInformation

	author := m.Authority{
		Name: book.Mods.Name.NamePart,
		Type: book.Mods.Name.Type,
		Role: book.Mods.Name.Role.RoleTerm.Text,
	}

	result.ArchiveID = archive.ArchiveID
	result.OaiIdentifier = archive.OaiIdentifier
	result.Platform = archive.Platform
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

func (ha *DetailArchive) mappingEprintsResponse(doc m.Archive) m.DetailEprintsResponse {
	var result m.DetailEprintsResponse
	var documents []m.Document

	result.ArchiveID = doc.ArchiveID
	result.OaiIdentifier = doc.OaiIdentifier
	result.Platform = doc.Platform
	result.Title = doc.Title
	result.Creator = doc.Creator
	result.Subject = doc.Subject
	result.Description = doc.Description
	result.Publisher = doc.Publisher
	result.Contributor = doc.Contributor
	result.Date = doc.Date
	result.Type = doc.Type
	result.Rights = doc.Rights

	identifiers := strings.Split(doc.Identifier, "|")
	languages := strings.Split(doc.Language, "|")
	formats := strings.Split(doc.Format, "|")

	for i := 0; i < len(formats); i++ {
		fn := strings.Split(identifiers[i], "/")
		d := m.Document{
			FileName: fn[len(fn)-1],
			Language: languages[i],
			Format:   formats[i],
			URL:      identifiers[i],
		}

		documents = append(documents, d)
	}

	result.DocumentIdentifier = identifiers[len(identifiers)-2]
	result.Documents = documents
	return result
}

func (ha *DetailArchive) mappingOJSResponse(doc m.Archive) m.Archive {

	relations := strings.Split(doc.Relation, "|")

	if len(relations) > 0 {
		doc.Relation = relations[0]
	}

	return doc

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
