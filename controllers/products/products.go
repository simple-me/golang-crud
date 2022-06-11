package user

import (
	"CRUD-Operation/db/conn"
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
	//fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	db := conn.GetPostgres()
	create := db.Create(&product.Product{Name: req.Name, Code: req.Code, Price: uint64(req.Price)})
	if create.Error != nil {
		c.JSON(http.StatusBadRequest, create.Error.Error())
	}
}

func FindProduct(c *gin.Context) {
	fmt.Println(c.Param("code"))
	prod := product.Product{}
	db := conn.GetPostgres()
	res := db.First(&prod, "code=?", c.Param("code"))
	if res.Error == nil {
		c.JSON(http.StatusOK, gin.H{"response": prod})
	}
}

func ListProducts(c *gin.Context) {
	var prod []product.Product
	db := conn.GetPostgres()
	res := db.Find(&prod)
	fmt.Println(res.RowsAffected)
	if res.Error == nil {
		c.JSON(http.StatusOK, gin.H{"response": prod})
	}
}

func UpdateProduct(c *gin.Context) {
	var req ProductParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	prod := product.Product{}
	db := conn.GetPostgres()
	err := db.First(&prod, "code=?", req.Code).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	prod.Price = uint64(req.Price)
	prod.Code = req.Code
	prod.Name = req.Name

	db.Save(&prod)
	c.JSON(http.StatusOK, "records changed")
}

func DeleteProduct(c *gin.Context) {
	var req ProductParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	prod := product.Product{}
	db := conn.GetPostgres()
	err := db.First(&prod, "code=?", req.Code).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	err = db.Delete(&prod).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

}
