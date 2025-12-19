package routes

import (
	"net/http"

	datasetApiSdk "github.com/ONSdigital/dp-dataset-api/sdk"

	ds "github.com/ONSdigital/dp-api-clients-go/v2/dataset"
	zc "github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	bc "github.com/ONSdigital/dp-publishing-dataset-controller/clients/topics"
	"github.com/ONSdigital/dp-publishing-dataset-controller/config"
	"github.com/ONSdigital/dp-publishing-dataset-controller/dataset"
	"github.com/gorilla/mux"
)

// Init initialises routes for the service
func Init(router *mux.Router, cfg *config.Config, hc healthcheck.HealthCheck, dc *ds.Client, zc *zc.Client, bc *bc.Client, datasetApiClient *datasetApiSdk.Client) {
	router.StrictSlash(true).Path("/health").HandlerFunc(hc.Handler)

	router.StrictSlash(true).Path("/datasets").HandlerFunc(dataset.GetAll(datasetApiClient, cfg.DatasetsBatchSize, cfg.DatasetsBatchWorkers)).Methods(http.MethodGet)
	//router.StrictSlash(true).Path("/datasets").HandlerFunc(dataset.GetAll(dc, cfg.DatasetsBatchSize, cfg.DatasetsBatchWorkers)).Methods(http.MethodGet)
	router.StrictSlash(true).Path("/datasets/{datasetID}/create").HandlerFunc(dataset.GetTopics(bc)).Methods(http.MethodGet)
	//router.StrictSlash(true).Path("/datasets/{datasetID}/editions").HandlerFunc(dataset.GetEditions(dc)).Methods(http.MethodGet)
	router.StrictSlash(true).Path("/datasets/{datasetID}/editions").HandlerFunc(dataset.GetEditions(datasetApiClient)).Methods(http.MethodGet)
	//router.StrictSlash(true).Path("/datasets/{datasetID}/editions/{editionID}/versions").HandlerFunc(dataset.GetVersions(dc, cfg.DatasetsBatchSize, cfg.DatasetsBatchWorkers)).Methods(http.MethodGet)
	router.StrictSlash(true).Path("/datasets/{datasetID}/editions/{editionID}/versions").HandlerFunc(dataset.GetVersions(datasetApiClient, cfg.DatasetsBatchSize, cfg.DatasetsBatchWorkers)).Methods(http.MethodGet)

	//router.StrictSlash(true).Path("/datasets/{datasetID}/editions/{editionID}/versions/{versionID}").HandlerFunc(dataset.GetMetadataHandler(dc, zc)).Methods(http.MethodGet)
	router.StrictSlash(true).Path("/datasets/{datasetID}/editions/{editionID}/versions/{versionID}").HandlerFunc(dataset.GetMetadataHandler(datasetApiClient, zc)).Methods(http.MethodGet)

	router.StrictSlash(true).Path("/datasets/{datasetID}/editions/{editionID}/versions/{versionID}").HandlerFunc(dataset.PutMetadata(datasetApiClient, zc)).Methods(http.MethodPut)
	router.StrictSlash(true).Path("/datasets/{datasetID}/editions/{editionID}/versions/{versionID}/metadata").HandlerFunc(dataset.PutEditableMetadata(datasetApiClient, zc)).Methods(http.MethodPut)
	//router.StrictSlash(true).Path("/datasets/{datasetID}/editions/{editionID}/versions/{versionID}/metadata").HandlerFunc(dataset.PutEditableMetadata(dc, zc)).Methods(http.MethodPut)
}
