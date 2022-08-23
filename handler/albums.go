package handler

import (
	"encoding/json"
	"fmt"
	"github.com/apod/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var URL = os.Getenv("APOD_API_URL") + os.Getenv("APOD_API_KEY")

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
	resp, err := http.Get(URL)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "response error"})
		return
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to NASA struct
	err = json.Unmarshal(bodyBytes, &nasa)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "convert error"})
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

	ctx.JSON(http.StatusOK, &model.ListNasa{Data: images})
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
	Url2 := URL + dateUrl
	resp, err := http.Get(Url2)
	if err != nil {
		logrus.Warnf("something went wrong")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "wrong response from API"})
		return
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &nasa)
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
	queryUrl := fmt.Sprintf("%s&thumbs=%t", URL, queryInput.Thumbs)

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

	body, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "You have exceeded your rate limit.") {
		logrus.Warnf("Error:%s", fmt.Errorf("you have exceeded your rate limit"))
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "something went wrong"})
		return
	}

	if !queryInput.Date.IsZero() {
		var outputSingle model.Nasa
		err = json.Unmarshal(body, &outputSingle)
		queryOutput = append(queryOutput, outputSingle)
	} else {
		err = json.Unmarshal(body, &queryOutput)
	}

	ctx.JSON(http.StatusOK, &model.ListNasa{Data: queryOutput})
}
