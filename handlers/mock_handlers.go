// Code generated by MockGen. DO NOT EDIT.
// Source: clients.go

// Package handlers is a generated GoMock package.
package handlers

import (
	context "context"
	dataset "github.com/ONSdigital/dp-api-clients-go/dataset"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDatasetClient is a mock of DatasetClient interface
type MockDatasetClient struct {
	ctrl     *gomock.Controller
	recorder *MockDatasetClientMockRecorder
}

// MockDatasetClientMockRecorder is the mock recorder for MockDatasetClient
type MockDatasetClientMockRecorder struct {
	mock *MockDatasetClient
}

// NewMockDatasetClient creates a new mock instance
func NewMockDatasetClient(ctrl *gomock.Controller) *MockDatasetClient {
	mock := &MockDatasetClient{ctrl: ctrl}
	mock.recorder = &MockDatasetClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatasetClient) EXPECT() *MockDatasetClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockDatasetClient) Get(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, datasetID string) (dataset.DatasetDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userAuthToken, serviceAuthToken, collectionID, datasetID)
	ret0, _ := ret[0].(dataset.DatasetDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDatasetClientMockRecorder) Get(ctx, userAuthToken, serviceAuthToken, collectionID, datasetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatasetClient)(nil).Get), ctx, userAuthToken, serviceAuthToken, collectionID, datasetID)
}

// GetByPath mocks base method
func (m *MockDatasetClient) GetByPath(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, path string) (dataset.DatasetDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPath", ctx, userAuthToken, serviceAuthToken, collectionID, path)
	ret0, _ := ret[0].(dataset.DatasetDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPath indicates an expected call of GetByPath
func (mr *MockDatasetClientMockRecorder) GetByPath(ctx, userAuthToken, serviceAuthToken, collectionID, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPath", reflect.TypeOf((*MockDatasetClient)(nil).GetByPath), ctx, userAuthToken, serviceAuthToken, collectionID, path)
}

// GetEditions mocks base method
func (m *MockDatasetClient) GetEditions(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, datasetID string) ([]dataset.Edition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEditions", ctx, userAuthToken, serviceAuthToken, collectionID, datasetID)
	ret0, _ := ret[0].([]dataset.Edition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEditions indicates an expected call of GetEditions
func (mr *MockDatasetClientMockRecorder) GetEditions(ctx, userAuthToken, serviceAuthToken, collectionID, datasetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEditions", reflect.TypeOf((*MockDatasetClient)(nil).GetEditions), ctx, userAuthToken, serviceAuthToken, collectionID, datasetID)
}

// GetEdition mocks base method
func (m *MockDatasetClient) GetEdition(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, datasetID, edition string) (dataset.Edition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEdition", ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition)
	ret0, _ := ret[0].(dataset.Edition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEdition indicates an expected call of GetEdition
func (mr *MockDatasetClientMockRecorder) GetEdition(ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEdition", reflect.TypeOf((*MockDatasetClient)(nil).GetEdition), ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition)
}

// GetVersions mocks base method
func (m *MockDatasetClient) GetVersions(ctx context.Context, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition string) ([]dataset.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersions", ctx, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition)
	ret0, _ := ret[0].([]dataset.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersions indicates an expected call of GetVersions
func (mr *MockDatasetClientMockRecorder) GetVersions(ctx, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersions", reflect.TypeOf((*MockDatasetClient)(nil).GetVersions), ctx, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition)
}

// GetVersion mocks base method
func (m *MockDatasetClient) GetVersion(ctx context.Context, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition, version string) (dataset.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", ctx, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition, version)
	ret0, _ := ret[0].(dataset.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockDatasetClientMockRecorder) GetVersion(ctx, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockDatasetClient)(nil).GetVersion), ctx, userAuthToken, serviceAuthToken, downloadServiceAuthToken, collectionID, datasetID, edition, version)
}

// GetVersionMetadata mocks base method
func (m *MockDatasetClient) GetVersionMetadata(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, id, edition, version string) (dataset.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionMetadata", ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version)
	ret0, _ := ret[0].(dataset.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionMetadata indicates an expected call of GetVersionMetadata
func (mr *MockDatasetClientMockRecorder) GetVersionMetadata(ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionMetadata", reflect.TypeOf((*MockDatasetClient)(nil).GetVersionMetadata), ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version)
}

// GetDimensions mocks base method
func (m *MockDatasetClient) GetDimensions(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, id, edition, version string) (dataset.Dimensions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensions", ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version)
	ret0, _ := ret[0].(dataset.Dimensions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensions indicates an expected call of GetDimensions
func (mr *MockDatasetClientMockRecorder) GetDimensions(ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockDatasetClient)(nil).GetDimensions), ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version)
}

// GetOptions mocks base method
func (m *MockDatasetClient) GetOptions(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, id, edition, version, dimension string) (dataset.Options, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOptions", ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version, dimension)
	ret0, _ := ret[0].(dataset.Options)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOptions indicates an expected call of GetOptions
func (mr *MockDatasetClientMockRecorder) GetOptions(ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version, dimension interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOptions", reflect.TypeOf((*MockDatasetClient)(nil).GetOptions), ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version, dimension)
}
