package handlers

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"github.com/raexera/soko/internal/app/models"
	"github.com/raexera/soko/internal/app/repositories"
)

func CreateProduct(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	newProduct, err := repositories.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newProduct)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := repositories.GetProductByID(idInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, product)
}

func GetAllProducts(c echo.Context) error {
	products, err := repositories.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	product := models.Product{}
	c.Bind(&product)
	updatedUser, err := repositories.UpdateProduct(product, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repositories.DeleteProduct(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
