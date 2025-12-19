package dataset

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	dpDatasetApiModels "github.com/ONSdigital/dp-dataset-api/models"
	datasetApiSdk "github.com/ONSdigital/dp-dataset-api/sdk"
	"github.com/gorilla/mux"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitGetAllDatasets(t *testing.T) {

	datasetsBatchSize := 10
	datasetsMaxWorkers := 3

	mockedDatasetResponse := []dpDatasetApiModels.DatasetUpdate{
		{
			ID: "id-1",
			Next: &dpDatasetApiModels.Dataset{
				Title: "Test title 1",
			},
		},
		{
			ID: "id-2",
			Next: &dpDatasetApiModels.Dataset{
				Title: "Test title 2",
			},
		},
	}

	expectedSuccessResponse := "[{\"id\":\"id-1\",\"title\":\"Test title 1\"},{\"id\":\"id-2\",\"title\":\"Test title 2\"}]"

	Convey("test getAllDatasets", t, func() {
		Convey("on success", func() {

			mockDatasetClient := &DatasetAPIClientMock{
				GetDatasetsInBatchesFunc: func(ctx context.Context, headers datasetApiSdk.Headers, batchSize int, maxWorkers int) (datasetApiSdk.DatasetsList, error) {
					return datasetApiSdk.DatasetsList{Items: mockedDatasetResponse}, nil
				},
			}

			req := httptest.NewRequest("GET", "/datasets", nil)
			req.Header.Set("Collection-Id", "testcollection")
			req.Header.Set("X-Florence-Token", "testuser")
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Path("/datasets").HandlerFunc(GetAll(mockDatasetClient, datasetsBatchSize, datasetsMaxWorkers))

			Convey("returns 200 response", func() {
				router.ServeHTTP(rec, req)
				So(rec.Code, ShouldEqual, http.StatusOK)
			})

			Convey("returns JSON response", func() {
				router.ServeHTTP(rec, req)
				response := rec.Body.String()
				So(response, ShouldEqual, expectedSuccessResponse)
			})
		})

		Convey("errors if no headers are passed", func() {

			mockDatasetClient := &DatasetAPIClientMock{
				GetDatasetsInBatchesFunc: func(ctx context.Context, headers datasetApiSdk.Headers, batchSize int, maxWorkers int) (datasetApiSdk.DatasetsList, error) {
					return datasetApiSdk.DatasetsList{}, nil
				},
			}

			Convey("collection id not set", func() {
				req := httptest.NewRequest("GET", "/datasets", nil)
				req.Header.Set("X-Florence-Token", "testuser")
				rec := httptest.NewRecorder()
				router := mux.NewRouter()
				router.Path("/datasets").HandlerFunc(GetAll(mockDatasetClient, datasetsBatchSize, datasetsMaxWorkers))

				Convey("returns 400 response", func() {
					router.ServeHTTP(rec, req)
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				})

				Convey("returns error body", func() {
					router.ServeHTTP(rec, req)
					response := rec.Body.String()
					So(response, ShouldResemble, "no collection ID header set\n")
				})
			})

			Convey("user auth token not set", func() {
				req := httptest.NewRequest("GET", "/datasets", nil)
				req.Header.Set("Collection-Id", "testcollection")
				rec := httptest.NewRecorder()
				router := mux.NewRouter()
				router.Path("/datasets").HandlerFunc(GetAll(mockDatasetClient, datasetsBatchSize, datasetsMaxWorkers))

				Convey("returns 400 response", func() {
					router.ServeHTTP(rec, req)
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				})

				Convey("returns error body", func() {
					router.ServeHTTP(rec, req)
					response := rec.Body.String()
					So(response, ShouldResemble, "no user access token header set\n")
				})
			})
		})

		Convey("handles error from dataset client", func() {

			mockDatasetClient := &DatasetAPIClientMock{
				GetDatasetsInBatchesFunc: func(ctx context.Context, headers datasetApiSdk.Headers, batchSize int, maxWorkers int) (datasetApiSdk.DatasetsList, error) {
					return datasetApiSdk.DatasetsList{}, errors.New("test dataset API error")
				},
			}

			req := httptest.NewRequest("GET", "/datasets", nil)
			req.Header.Set("Collection-Id", "testcollection")
			req.Header.Set("X-Florence-Token", "testuser")
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Path("/datasets").HandlerFunc(GetAll(mockDatasetClient, datasetsBatchSize, datasetsMaxWorkers))

			Convey("returns 500 response", func() {
				router.ServeHTTP(rec, req)
				So(rec.Code, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("returns error body", func() {
				router.ServeHTTP(rec, req)
				response := rec.Body.String()
				So(response, ShouldResemble, "error getting all datasets from dataset API\n")
			})

		})
	})
}
