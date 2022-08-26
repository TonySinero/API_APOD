package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apod/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

//	swagger:route POST /album/ album createAlbum
//
//	Create album
//
//	Creates a new album.
//
//	responses:
//	 201: map[string]int
//	 400: model.ErrorResponse
//	 500: model.ErrorResponse
func (h *Handler) createAlbum(ctx *gin.Context) {
	var nasa model.Nasa
	resp, err := http.Get(os.Getenv("APOD_API_URL") + os.Getenv("APOD_API_KEY"))
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "http response error"})
		return
	}

	defer resp.Body.Close()

	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "io.copy error"})
		return
	}

	// Convert response body to NASA struct
	err = json.Unmarshal(buf.Bytes(), &nasa)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "unmarshal error"})
		return
	}
	Id, err := h.services.Image.CreateAlbum(&nasa)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]int{
		"id": Id,
	})
}

//	swagger:route GET /album/images album getAlbumFromDB
//
//	Get album from database.
//
//	Returns an albums.
//
//	responses:
//	 200: model.ListNasa
//	 401: model.ErrorResponse
//	 500: model.ErrorResponse
func (h *Handler) getAlbumFromDB(ctx *gin.Context) {
	images, err := h.services.Image.GetAlbumFromDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, images)
}

//	swagger:route GET /album/dt album getByDate
//
//	Get album by date.
//
//	Returns an album with the given date.
//
//	responses:
//	 200: model.Nasa
//	 401: model.ErrorResponse
//	 500: model.ErrorResponse
func (h *Handler) getByDate(ctx *gin.Context) {
	var nasa model.Nasa
	dateParam := ctx.Query("date")
	dateUrl := fmt.Sprintf("&date=%s", dateParam)
	Url2 := os.Getenv("APOD_API_URL") + os.Getenv("APOD_API_KEY") + dateUrl
	resp, err := http.Get(Url2)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "wrong response from API"})
		return
	}

	defer resp.Body.Close()

	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "io.copy error"})
		return
	}

	err = json.Unmarshal(buf.Bytes(), &nasa)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "unmarshall error"})
		return
	}
	ctx.JSON(http.StatusOK, nasa)
}

//	swagger:route GET /album/filter album getWithFilter
//
//	Get album by filter.
//
//	Returns an albums with the given filter.
//
//	responses:
//	 200: model.ListNasa
//	 401: model.ErrorResponse
//	 500: model.ErrorResponse
func (h *Handler) getWithFilter(ctx *gin.Context) {
	var queryOutput []model.Nasa
	var queryInput model.ApodQueryInput
	queryUrl := fmt.Sprintf("%s&thumbs=%t", os.Getenv("APOD_API_URL")+os.Getenv("APOD_API_KEY"), queryInput.Thumbs)

	if queryInput.Count != 0 {
		if !queryInput.Date.IsZero() || !queryInput.StartDate.IsZero() || !queryInput.EndDate.IsZero() {
			logrus.Warnf("Error:%s", fmt.Errorf("cannot use the following params with 'Count': 'Date', 'StartDate', 'EndDate'"))
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "something went wrong"})
			return
		}

		queryUrl += fmt.Sprintf("&count=%d", queryInput.Count)
	}

	if !queryInput.Date.IsZero() {
		if !queryInput.StartDate.IsZero() || !queryInput.EndDate.IsZero() {
			logrus.Warnf("Error:%s", fmt.Errorf("cannot use params 'Date' and 'StartDate' or 'EndDate' together"))
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "something went wrong"})
			return
		}

		queryUrl += fmt.Sprintf("&date=%s", queryInput.Date.Format("2006-01-02"))
	}

	if !queryInput.StartDate.IsZero() {
		queryUrl += fmt.Sprintf("&start_date=%s", queryInput.StartDate.Format("2006-01-02"))
	}

	if !queryInput.EndDate.IsZero() {
		queryUrl += fmt.Sprintf("&end_date=%s", queryInput.EndDate.Format("2006-01-02"))
	}

	resp, err := http.Get(queryUrl)
	if err != nil {
		logrus.Warnf("Wrong response")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "something went wrong"})
		return
	}
	defer resp.Body.Close()

	if resp.Header.Get("X-RateLimit-Remaining") == "0" {
		logrus.Warnf("Error:%s", fmt.Errorf("you have exceeded your rate limit"))
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "something went wrong"})
		return
	}
	buf := bytes.NewBuffer(nil)

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "io.copy error"})
		return
	}

	if !queryInput.Date.IsZero() {
		var outputSingle model.Nasa
		err = json.Unmarshal(buf.Bytes(), &outputSingle)
		queryOutput = append(queryOutput, outputSingle)
	} else {
		err = json.Unmarshal(buf.Bytes(), &queryOutput)
	}

	ctx.JSON(http.StatusOK, &model.ListNasa{Data: queryOutput})
}
