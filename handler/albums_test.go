package handler

import (
	"bytes"
	"fmt"
	"github.com/apod/model"
	"github.com/apod/service"
	mockservice "github.com/apod/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetAlbumFromDB(t *testing.T) {
	type mockBehavior func(s *mockservice.MockImage)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mockservice.MockImage) {
				s.EXPECT().GetAlbumFromDB().Return([]model.Nasa{
					{
						ID:             1,
						Copyright:      "test juan",
						Date:           "2022-08-18",
						Explanation:    "tt",
						Hdurl:          "ttt",
						MediaType:      "image",
						ServiceVersion: "test v1",
						Title:          "test title",
						URL:            "test url",
					},
					{
						ID:             2,
						Copyright:      "test juan2",
						Date:           "2022-08-19",
						Explanation:    "tt2",
						Hdurl:          "ttt2",
						MediaType:      "image2",
						ServiceVersion: "test v2",
						Title:          "test title2",
						URL:            "test url2",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"id":1,"copyright":"test juan","date":"2022-08-18","explanation":"tt","hdurl":"ttt","media_type":"image","service_version":"test v1","title":"test title","url":"test url"},{"id":2,"copyright":"test juan2","date":"2022-08-19","explanation":"tt2","hdurl":"ttt2","media_type":"image2","service_version":"test v2","title":"test title2","url":"test url2"}]`,
		},
		{
			name: "Server error",
			mockBehavior: func(s *mockservice.MockImage) {
				s.EXPECT().GetAlbumFromDB().Return(nil, fmt.Errorf("server error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			appService := mockservice.NewMockImage(c)
			testCase.mockBehavior(appService)
			serv := &service.Service{Image: appService}
			handler := NewHandler(serv)

			r := handler.InitRoutes()

			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", fmt.Sprintf("/album/images"), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestCreateAlbum(t *testing.T) {
	type mockBehavior func(s *mockservice.MockImage, image *model.Nasa)

	testTable := []struct {
		name                string
		inputBody           string
		inputNasa           *model.Nasa
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "Incorrect data came from the request",
			inputBody:           `{"id":1,"copyright":"test juan","date":"2022-08-18","explanation":"tt","hdurl":"ttt","media_type":"image","service_version":"test v1","title":"test title","url":"test url"}`,
			inputNasa:           &model.Nasa{},
			mockBehavior:        func(s *mockservice.MockImage, image *model.Nasa) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"http response error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			appService := mockservice.NewMockImage(c)
			testCase.mockBehavior(appService, testCase.inputNasa)
			serv := &service.Service{Image: appService}
			handler := NewHandler(serv)

			r := handler.InitRoutes()

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/album/", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
