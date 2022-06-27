package user

import (
	model "CRUD-Operation/models/products"
	product "CRUD-Operation/models/products"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductParams struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Price int64  `json:"price"`
}

func CreateProduct(c *gin.Context) {
	var req ProductParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := model.Create(product.Product{Name: req.Name, Code: req.Code, Price: uint64(req.Price)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}

func FindProduct(c *gin.Context) {
	fmt.Println(c.Param("code"))
	prod, err := model.Get(c.Param("code"))
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": prod})
}

func ListProducts(c *gin.Context) {
	prod, err := model.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": prod})
}

func UpdateProduct(c *gin.Context) {
	var req ProductParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := model.Update(product.Product{Name: req.Name, Code: req.Code, Price: uint64(req.Price)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}

func DeleteProduct(c *gin.Context) {
	var req ProductParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := model.Delete(req.Code)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, "record deleted")
}
