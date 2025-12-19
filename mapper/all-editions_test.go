package mapper

import (
	"testing"

	"github.com/ONSdigital/dp-dataset-api/models"
	datasetApiSdk "github.com/ONSdigital/dp-dataset-api/sdk"
	"github.com/ONSdigital/dp-publishing-dataset-controller/model"

	. "github.com/smartystreets/goconvey/convey"
)

var mockedDataset = models.DatasetUpdate{
	Next: &models.Dataset{
		Title: "Test title",
	},
}

var mockedEditions = datasetApiSdk.EditionsList{
	Items: []models.Edition{models.Edition{Edition: "edition-1"}, models.Edition{Edition: "edition-2"}},
}

var mockedLatestVersions = map[string]string{"edition-1": "2020-11-07T00:00:00.000Z", "edition-2": ""}

func TestUnitAllEditions(t *testing.T) {
	t.Parallel()

	expectedAllEditions := []model.Edition{{ID: "edition-1", Title: "edition-1", ReleaseDate: "07 November 2020"}, {ID: "edition-2", Title: "edition-2", ReleaseDate: ""}}

	expectedEditionsPage := model.EditionsPage{DatasetName: "Test title", Editions: expectedAllEditions}

	Convey("test all editions maps correctly", t, func() {
		mapped := AllEditions(ctx, mockedDataset, mockedEditions, mockedLatestVersions)
		So(mapped, ShouldResemble, expectedEditionsPage)
	})
}
