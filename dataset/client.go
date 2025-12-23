package dataset

import (
	"context"

	zebedeeclient "github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	datasetApiModels "github.com/ONSdigital/dp-dataset-api/models"
	datasetApiSdk "github.com/ONSdigital/dp-dataset-api/sdk"
	babbageclient "github.com/ONSdigital/dp-publishing-dataset-controller/clients/topics"
)

//go:generate moq -out mocks_test.go -pkg dataset . DatasetAPIClient ZebedeeClient BabbageClient

type DatasetAPIClient interface {
	GetDatasetsInBatches(ctx context.Context, headers datasetApiSdk.Headers, batchSize, maxWorkers int) (datasetApiSdk.DatasetsList, error)
	GetEdition(ctx context.Context, headers datasetApiSdk.Headers, datasetID, edition string) (datasetApiModels.Edition, error)
	GetEditions(ctx context.Context, headers datasetApiSdk.Headers, datasetID string, q *datasetApiSdk.QueryParams) (m datasetApiSdk.EditionsList, err error)
	GetDatasetCurrentAndNext(ctx context.Context, headers datasetApiSdk.Headers, datasetID string) (m datasetApiModels.DatasetUpdate, err error)
	GetVersionWithHeaders(ctx context.Context, headers datasetApiSdk.Headers, datasetID, edition, version string) (v datasetApiModels.Version, h datasetApiSdk.ResponseHeaders, err error)
	GetVersion(ctx context.Context, headers datasetApiSdk.Headers, datasetID, edition, version string) (m datasetApiModels.Version, err error)
	GetVersionsInBatches(ctx context.Context, headers datasetApiSdk.Headers, datasetID, edition string, batchSize, maxWorkers int) (versions datasetApiSdk.VersionsList, err error)
	PutDataset(ctx context.Context, headers datasetApiSdk.Headers, datasetID string, d datasetApiModels.Dataset) error
	PutMetadata(ctx context.Context, headers datasetApiSdk.Headers, datasetID, edition, version string, metadata datasetApiModels.EditableMetadata, versionEtag string) error
	PutVersion(ctx context.Context, headers datasetApiSdk.Headers, datasetID, editionID, versionID string, version datasetApiModels.Version) (updatedVersion datasetApiModels.Version, err error)
	PutInstance(ctx context.Context, headers datasetApiSdk.Headers, instanceID string, i datasetApiSdk.UpdateInstance, ifMatch string) (eTag string, err error)
}

type ZebedeeClient interface {
	GetCollection(ctx context.Context, userAccessToken, collectionID string) (c zebedeeclient.Collection, err error)
	PutDatasetInCollection(ctx context.Context, userAccessToken, collectionID, lang, datasetID, state string) error
	PutDatasetVersionInCollection(ctx context.Context, userAccessToken, collectionID, lang, datasetID, edition, version, state string) error
}

type BabbageClient interface {
	GetTopics(ctx context.Context, userAccessToken string) (result babbageclient.TopicsResult, err error)
}
