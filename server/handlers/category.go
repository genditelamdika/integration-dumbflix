package handlers

import (
	categoriesdto "dumbmerch/dto/category"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
}

func HandlerCategory(CategoryRepository repositories.CategoryRepository) *handlerCategory {
	return &handlerCategory{CategoryRepository}
}
func (h *handlerCategory) FindCategory(c echo.Context) error {
	categories, err := h.CategoryRepository.FindCategory()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: categories})
}
func (h *handlerCategory) GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// responseCategory := categoriesdto.CategoryResponse{
	// 	ID:   category.ID,
	// 	Name: category.Name,
	// 	Film: []categoriesdto.CategoryFilm{},
	// }

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsecategory(category)})
}
func (h *handlerCategory) CreateCategory(c echo.Context) error {
	request := new(categoriesdto.CreateCategoryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	category := models.Category{
		Name: request.Name,
		// Films: request.Film,
		// FilmID: request.FilmID,
		// Films: request.Films,
		// Email:    request.Email,
		// Password: request.Password,
	}

	data, err := h.CategoryRepository.CreateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponsecategory(data)})
}
func (h *handlerCategory) UpdateCategory(c echo.Context) error {
	request := new(categoriesdto.UpdateCategoryRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	category, err := h.CategoryRepository.GetCategory(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		category.Name = request.Name
	}

	// if request.Email != "" {
	// 	user.Email = request.Email
	// }

	// if request.Password != "" {
	// 	user.Password = request.Password
	// }

	data, err := h.CategoryRepository.UpdateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
func (h *handlerCategory) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CategoryRepository.DeleteCategory(category, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

// func mapFilmToDto(u []models.Film) (mapFilmResponse []models.CategoryFilm) {
// 	// for _, filmModel := range u {
// 		filmResponse := models.CategoryFilm{
// 			ID:            u.ID,
// 			Title:         u.Title,
// 			Description:   u.Description,
// 			Year:          u.Year,
// 			ThumbnailFilm: "http://localhost:5000/uploads/" + filmModel.ThumbnailFilm,
// 		}
// 		mapFilmResponse = append(mapFilmResponse, filmResponse)
// 	}
// return
// }

func convertResponsecategory(u models.Category) models.Category {
	return models.Category{
		ID:   u.ID,
		Name: u.Name,
		// ThumbnailFilm: "http://localhost:5000/uploads/" + u.ThumbnailFilm,
		// Film: u.Film[],
		// FilmID: u.FilmID,
		// Films:  u.Films,
		// Email:    u.Email,
		// Password: u.Password,
	}
}
