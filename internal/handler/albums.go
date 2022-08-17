package handler

import (
	"encoding/json"
	"fmt"
	"github.com/apod/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// @Summary createPerm
// @Tags images
// @Description create new albums
// @Accept  json
// @Produce  json
// @Param input body model.Nasa true "NASA"
// @Success 201 {object} map[string]int
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /album/ [post]
func (h *Handler) createAlbum(ctx *gin.Context) {
	var nasa model.Nasa
	var URL = os.Getenv("APOD_API_URL") + os.Getenv("APOD_API_KEY")
	resp, err := http.Get(URL)
	if err != nil {
		logrus.Warnf("Wrong response")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "something went wrong"})
		return
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to NASA struct
	err = json.Unmarshal(bodyBytes, &nasa)
	if err != nil {
		logrus.Warnf("Wrong response")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
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

// @Summary getAllAlbums
// @Tags images
// @Description gets all albums from API
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ListNasa
// @Failure 500 {object} model.ErrorResponse
// @Router /album/ [get]
func (h *Handler) getAllAlbums(ctx *gin.Context) {
	var URL string
	var date string
	var count string
	if ctx.Query("date") != "" || ctx.Query("count") != "" {
		date = ctx.Query("date")
		count = ctx.Query("count")
		URL = os.Getenv("APOD_API_URL") + os.Getenv("APOD_API_KEY") + "&count=" + count + "&date=" + date
	} else {
		URL = os.Getenv("APOD_API_URL") + os.Getenv("APOD_API_KEY")
	}

	resp, err := http.Get(URL)
	if err != nil {
		logrus.Warnf("Wrong response")
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Something went wrong"})
		return
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var nasaArray = []model.Nasa{}
	err = json.Unmarshal(bodyBytes, &nasaArray)

	if len(nasaArray) == 0 {
		logrus.Warnf("Empty array:%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Empty array"})
		return
	}

	ctx.JSON(http.StatusOK, &model.ListNasa{Data: nasaArray})
}

// @Summary getImagesFromDB
// @Tags images
// @Description gets all images from DB
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ListNasa
// @Failure 500 {object} model.ErrorResponse
// @Router /album/images [get]
func (h *Handler) getAlbumFromDB(ctx *gin.Context) {
	images, err := h.services.Image.GetAlbumFromDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &model.ListNasa{Data: images})
}

// @Summary getWithFilter
// @Tags images
// @Description gets all images with filter
// @Accept  json
// @Produce  json
// @Param input body model.ApodQueryInput true "Filter"
// @Success 200 {object} model.ListNasa
// @Failure 500 {object} model.ErrorResponse
// @Router /album/filter [get]
func (h *Handler) getWithFilter(ctx *gin.Context) {
	var queryOutput []model.Nasa
	var queryInput model.ApodQueryInput
	URL := os.Getenv("APOD_API_URL") + os.Getenv("APOD_API_KEY")
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
