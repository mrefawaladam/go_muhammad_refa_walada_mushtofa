package handler

import (
	"log"
	"net/http"
	"strconv"

	"capston-lms/internal/application/usecase"
	"capston-lms/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ItemsHandler struct {
	ItemUseCase usecase.ItemUseCase
}

func (handler ItemsHandler) GetItems() echo.HandlerFunc {
	return func(e echo.Context) error {
		var items []entity.Items
		keyword := e.QueryParam("keyword")
		log.Printf("Incoming request:", keyword)

		if keyword != "" {
			// Jika keyword ada, gunakan keyword
			items, err := handler.ItemUseCase.GetItemByName(keyword)
			if err != nil {
				return e.JSON(500, echo.Map{
					"error": err.Error(),
				})
			}

			return e.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get item by name",
				"item":    items,
			})
		}

		// Jika keyword kosong, gunakan "all"
		items, err := handler.ItemUseCase.GetAlItem()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all items",
			"items":   items,
		})
	}
}

func (handler ItemsHandler) GetItem() echo.HandlerFunc {
	return func(e echo.Context) error {
		var item entity.Items
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		item, err = handler.ItemUseCase.GetItem(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get item",
			"item":    item,
		})
	}
}
func (handler ItemsHandler) GetItemByCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		var item entity.Items
		category_id, err := strconv.Atoi(e.Param("category_id"))
		log.Printf("Incoming request:", category_id)

		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		item, err = handler.ItemUseCase.GetItemByCategory(category_id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get item",
			"item":    item,
		})
	}
}

func (handler ItemsHandler) CreateItem() echo.HandlerFunc {
	return func(e echo.Context) error {
		var item entity.Items
		if err := e.Bind(&item); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(item); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}

		err := handler.ItemUseCase.CreateItem(item)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create item"})
		}

		return e.JSON(http.StatusCreated, item)
	}
}

func (handler ItemsHandler) UpdateItem() echo.HandlerFunc {
	var item entity.Items

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.ItemUseCase.FindItem(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&item); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.ItemUseCase.UpdateItem(id, item)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update item",
		})
	}
}

func (handler ItemsHandler) DeleteItem() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.ItemUseCase.DeleteIetem(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete item`",
		})
	}
}
