package dataset

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	datasetApiModels "github.com/ONSdigital/dp-dataset-api/models"
	datasetApiSdk "github.com/ONSdigital/dp-dataset-api/sdk"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitGetVersions(t *testing.T) {
	t.Parallel()

	datasetID := "test-dataset"
	editionID := "test-edition"
	verionsBatchSize := 10
	versionsMaxWorkers := 3

	mockedDatasetResponse := datasetApiModels.DatasetUpdate{
		Next: &datasetApiModels.Dataset{
			Title: "Test title",
		},
	}

	mockedEditionResponse := datasetApiModels.Edition{
		Edition: "edition-1",
	}

	mockedVersionsResponse := []datasetApiModels.Version{
		{
			ID:      "version-1",
			Version: 1,
		},
		{
			ID:      "version-2",
			Version: 2,
		},
	}

	expectedSuccessResponse := "{\"dataset_name\":\"Test title\",\"edition_name\":\"edition-1\",\"versions\":[{\"id\":\"version-2\",\"title\":\"Version: 2\",\"version\":2,\"release_date\":\"\",\"state\":\"\"},{\"id\":\"version-1\",\"title\":\"Version: 1\",\"version\":1,\"release_date\":\"\",\"state\":\"\"}]}"

	Convey("test getAllVersions", t, func() {
		mockDatasetClient := &DatasetAPIClientMock{
			GetDatasetCurrentAndNextFunc: func(ctx context.Context, headers datasetApiSdk.Headers, datasetID string) (datasetApiModels.DatasetUpdate, error) {
				return mockedDatasetResponse, nil
			},
			GetEditionFunc: func(ctx context.Context, headers datasetApiSdk.Headers, datasetID string, editionID string) (datasetApiModels.Edition, error) {
				return mockedEditionResponse, nil
			},
			GetVersionsInBatchesFunc: func(ctx context.Context, headers datasetApiSdk.Headers, datasetID string, editionID string, batchSize int, maxWorkers int) (datasetApiSdk.VersionsList, error) {
				return datasetApiSdk.VersionsList{Items: mockedVersionsResponse}, nil
			},
		}

		Convey("on success", func() {
			reqURL := fmt.Sprintf("/datasets/%v/editions/%v/versions", datasetID, editionID)
			req := httptest.NewRequest("GET", reqURL, http.NoBody)
			req.Header.Set("Collection-Id", "testcollection")
			req.Header.Set("X-Florence-Token", "testuser")
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Path(reqURL).HandlerFunc(GetVersions(mockDatasetClient, verionsBatchSize, versionsMaxWorkers))

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
			Convey("collection id not set", func() {
				reqURL := fmt.Sprintf("/datasets/%v/editions/%v/versions", datasetID, editionID)
				req := httptest.NewRequest("GET", reqURL, http.NoBody)
				req.Header.Set("X-Florence-Token", "testuser")
				rec := httptest.NewRecorder()
				router := mux.NewRouter()
				router.Path(reqURL).HandlerFunc(GetVersions(mockDatasetClient, verionsBatchSize, versionsMaxWorkers))

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
				reqURL := fmt.Sprintf("/datasets/%v/editions/%v/versions", datasetID, editionID)
				req := httptest.NewRequest("GET", reqURL, http.NoBody)
				req.Header.Set("Collection-Id", "testcollection")
				rec := httptest.NewRecorder()
				router := mux.NewRouter()
				router.Path(reqURL).HandlerFunc(GetVersions(mockDatasetClient, verionsBatchSize, versionsMaxWorkers))

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
				GetDatasetCurrentAndNextFunc: func(ctx context.Context, headers datasetApiSdk.Headers, datasetID string) (datasetApiModels.DatasetUpdate, error) {
					return mockedDatasetResponse, nil
				},
				GetEditionFunc: func(ctx context.Context, headers datasetApiSdk.Headers, datasetID string, editionID string) (datasetApiModels.Edition, error) {
					return mockedEditionResponse, nil
				},
				GetVersionsInBatchesFunc: func(ctx context.Context, headers datasetApiSdk.Headers, datasetID string, editionID string, batchSize int, maxWorkers int) (datasetApiSdk.VersionsList, error) {
					return datasetApiSdk.VersionsList{}, errors.New("test dataset API error")
				},
			}

			reqURL := fmt.Sprintf("/datasets/%v/editions/%v/versions", datasetID, editionID)
			req := httptest.NewRequest("GET", reqURL, http.NoBody)
			req.Header.Set("Collection-Id", "testcollection")
			req.Header.Set("X-Florence-Token", "testuser")
			rec := httptest.NewRecorder()
			router := mux.NewRouter()
			router.Path(reqURL).HandlerFunc(GetVersions(mockDatasetClient, verionsBatchSize, versionsMaxWorkers))

			Convey("returns 500 response", func() {
				router.ServeHTTP(rec, req)
				So(rec.Code, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("returns error body", func() {
				router.ServeHTTP(rec, req)
				response := rec.Body.String()
				So(response, ShouldResemble, "error getting all versions from dataset API: test dataset API error\n")
			})
		})
	})
}
