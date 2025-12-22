package mapper

import (
	"context"
	"time"

	dpDatasetApiModels "github.com/ONSdigital/dp-dataset-api/models"
	datasetApiSdk "github.com/ONSdigital/dp-dataset-api/sdk"

	"github.com/ONSdigital/dp-publishing-dataset-controller/model"
	"github.com/ONSdigital/log.go/v2/log"
)

// AllEditions maps dataset and editions response to editions list page model
func AllEditions(ctx context.Context, dataset dpDatasetApiModels.DatasetUpdate, editions datasetApiSdk.EditionsList, latestVersions map[string]string) model.EditionsPage {
	mappedEditions := make([]model.Edition, len(editions.Items))
	for i := range editions.Items {
		var timeF string
		for k, latestVersion := range latestVersions {
			if k == editions.Items[i].Edition {
				timeParse, err := time.Parse("2006-01-02T15:04:05Z", latestVersion)
				if err != nil {
					log.Warn(ctx, "failed to parse release date", log.FormatErrors([]error{err}))
				} else {
					timeF = timeParse.Format("02 January 2006")
				}
			}
		}
		mappedEditions[i] = model.Edition{
			ID:          editions.Items[i].Edition,
			Title:       editions.Items[i].Edition,
			ReleaseDate: timeF,
		}
	}

	return model.EditionsPage{
		DatasetName: dataset.Next.Title,
		Editions:    mappedEditions,
	}
}
