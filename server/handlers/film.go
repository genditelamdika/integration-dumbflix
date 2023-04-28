package handlers

import (
	"context"
	filmsdto "dumbmerch/dto/film"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var path_file = "http://localhost:5000/uploads/"

type handlerfilm struct {
	FilmRepository repositories.FilmRepository
}

func HandlerFilm(FilmRepository repositories.FilmRepository) *handlerfilm {
	return &handlerfilm{FilmRepository}
}

func (h *handlerfilm) FindFilm(c echo.Context) error {
	films, err := h.FilmRepository.FindFilm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: films})
}

func (h *handlerfilm) GetFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsefilm(film)})
}
func (h *handlerfilm) CreateFilm(c echo.Context) error {
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	year, _ := strconv.Atoi(c.FormValue("year"))
	// qty, _ := strconv.Atoi(c.FormValue("qty"))
	categoryid, _ := strconv.Atoi(c.FormValue("categoryid"))

	request := filmsdto.CreateFilmRequest{
		Title:         c.FormValue("title"),
		ThumbnailFilm: dataFile,
		Year:          year,
		Linkfilm:      c.FormValue("linkfilm"),
		CategoryID:    categoryid,
		Description:   c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "uploads"})

	if err != nil {
		fmt.Println(err.Error())
	}

	// userLogin := c.Get("userLogin")
	// categoryId := userLogin.(jwt.MapClaims)["id"].(float64)

	film := models.Film{
		Title:         request.Title,
		ThumbnailFilm: resp.SecureURL,
		Year:          request.Year,
		CategoryID:    request.CategoryID,
		Description:   request.Description,
	}

	film, err = h.FilmRepository.CreateFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	film, _ = h.FilmRepository.GetFilm(film.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: film})
}

// func (h *handlerfilm) CreateFilm(c echo.Context) error {
// 	request := new(filmsdto.CreateFilmRequest)
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	film := models.Film{
// 		Title:         request.Title,
// 		ThumbnailFilm: request.ThumbnailFilm,
// 		Year:          request.Year,
// 		CategoryID:    request.CategoryID,
// 		Category:      request.Category,
// 		Description:   request.Description,
// 	}

// 	data, err := h.FilmRepository.CreateFilm(film)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsefilm(data)})
// }

func (h *handlerfilm) UpdateFilm(c echo.Context) error {
	// request := new(filmsdto.UpdateFilmRequest)
	// if err := c.Bind(&request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "uploads"})

	if err != nil {
		fmt.Println(err.Error())
	}

	year, _ := strconv.Atoi(c.FormValue("year"))
	// qty, _ := strconv.Atoi(c.FormValue("qty"))
	categoryid, _ := strconv.Atoi(c.FormValue("categoryid"))

	request := filmsdto.CreateFilmRequest{
		Title:         c.FormValue("title"),
		ThumbnailFilm: resp.SecureURL,
		Year:          year,
		// Linkfilm:      c.FormValue("linkfilm"),
		CategoryID:  categoryid,
		Description: c.FormValue("description"),
	}

	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		film.Title = request.Title
	}

	if request.ThumbnailFilm != "" {
		film.ThumbnailFilm = request.ThumbnailFilm
	}
	if request.Year != 0 {
		film.Year = request.Year
	}
	// if request.Linkfilm != "" {
	// 	film.Title = request.Linkfilm
	// }

	if request.CategoryID != 0 {
		film.CategoryID = request.CategoryID
	}
	// if request.Category.Name != "" {
	// 	film.Category = request.Category
	// }
	if request.Description != "" {
		film.Description = request.Description
	}
	dataCategory, _ := h.FilmRepository.GetCategoryfilm(film.CategoryID)
	// data, err := h.FilmRepository.UpdateFilm(user, id)

	data, err := h.FilmRepository.UpdateFilm(film, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	data.Category = dataCategory

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsefilm(data)})
}

func (h *handlerfilm) DeleteFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FilmRepository.DeleteFilm(film)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsefilm(data)})
}

func convertResponsefilm(u models.Film) models.Film {
	return models.Film{
		ID:            u.ID,
		Title:         u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		Year:          u.Year,
		Category:      u.Category,
		// CategoryID:    u.CategoryID,
		Description: u.Description,
	}
}
