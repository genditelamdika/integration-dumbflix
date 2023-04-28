package handlers

import (
	episodedto "dumbmerch/dto/episode"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// var path_file = "http://localhost:5000/uploads/"

type handlerepisode struct {
	EpisodeRepository repositories.EpisodeRepository
}

func HandlerEpisode(EpisodeRepository repositories.EpisodeRepository) *handlerepisode {
	return &handlerepisode{EpisodeRepository}
}

func (h *handlerepisode) FindEpisode(c echo.Context) error {
	episodes, err := h.EpisodeRepository.FindEpisode()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	for i, p := range episodes {
		episodes[i].ThumbnailEpisode = path_file + p.ThumbnailEpisode
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episodes})
}

func (h *handlerepisode) FindEpisodeByFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	episode, err := h.EpisodeRepository.FindEpisodeByFilm(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
}
func (h *handlerepisode) GetEpisodeByFilm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ide, _ := strconv.Atoi(c.Param("ide"))

	episode, err := h.EpisodeRepository.GetEpisodeByFilm(id, ide)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	episode.ThumbnailEpisode = path_file + episode.ThumbnailEpisode
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
}

func (h *handlerepisode) GetEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	episode.ThumbnailEpisode = path_file + episode.ThumbnailEpisode

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(episode)})
}

func (h *handlerepisode) CreateEpisode(c echo.Context) error {
	// get the datafile here
	dataFile := c.Get("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	filmid, _ := strconv.Atoi(c.FormValue("filmid"))

	request := episodedto.CreateEpisodeRequest{
		TitleEpisode:     c.FormValue("titleepisode"),
		ThumbnailEpisode: dataFile,
		LinkFilm:         c.FormValue("linkfilm"),
		FilmID:           filmid,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// userLogin := c.Get("userLogin")
	// categoryId := userLogin.(jwt.MapClaims)["id"].(float64)

	episode := models.Episode{
		TitleEpisode:     request.TitleEpisode,
		ThumbnailEpisode: request.ThumbnailEpisode,
		LinkFilm:         request.LinkFilm,
		FilmID:           request.FilmID,
		// Description:   request.Description,
	}

	episode, err = h.EpisodeRepository.CreateEpisode(episode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	episode, _ = h.EpisodeRepository.GetEpisode(episode.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: episode})
	// request := new(episodedto.CreateEpisodeRequest)
	// if err := c.Bind(request); err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// }

	// episode := models.Episode{
	// 	TitleEpisode:     request.TitleEpisode,
	// 	ThumbnailEpisode: request.ThumbnailEpisode,
	// 	LinkFilm:         request.LinkFilm,
	// 	FilmID:           request.FilmID,
	// 	Film:             request.Film,
	// }

	// data, err := h.EpisodeRepository.CreateEpisode(episode)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	// }

	// return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func (h *handlerepisode) UpdateEpisode(c echo.Context) error {
	request := new(episodedto.UpdateEpisodeRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.TitleEpisode != "" {
		episode.TitleEpisode = request.TitleEpisode
	}

	if request.ThumbnailEpisode != "" {
		episode.ThumbnailEpisode = request.ThumbnailEpisode
	}

	if request.LinkFilm != "" {
		episode.LinkFilm = request.LinkFilm
	}
	if request.FilmID != 0 {
		episode.FilmID = request.FilmID
	}
	if request.Film.ID != 0 {
		episode.Film = request.Film
	}

	data, err := h.EpisodeRepository.UpdateEpisode(episode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func (h *handlerepisode) DeleteEpisode(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.EpisodeRepository.DeleteEpisode(episode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)})
}

func convertResponseEpisode(u models.Episode) models.Episode {
	return models.Episode{
		ID:               u.ID,
		TitleEpisode:     u.TitleEpisode,
		ThumbnailEpisode: u.ThumbnailEpisode,
		LinkFilm:         u.LinkFilm,
		FilmID:           u.FilmID,
		// Film:          u.Film,
		// Category:      u.Category,
	}
}
