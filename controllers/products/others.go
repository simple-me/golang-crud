package user

import (
	"CRUD-Operation/db/conn"
	product "CRUD-Operation/models/products"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductParams struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Price int64  `json:"price"`
}

func CreateProduct(c *gin.Context) {
	var req ProductParams
	fmt.Println(c.Request.Body)
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
	//prod := product.Product{}
	db := conn.GetPostgres()
	//product_to_update := db.Find(&req, "code=?", req.Code)
	/* if product_to_update.Error != nil {
		c.JSON(http.StatusNotFound, product_to_update.Error.Error())
	} */
	db.Where("code=?", req.Code).First("products")
	//query_name := c.Request.URL.Query()
	//db.Where("code=?", query_name["code"]).First(&prod)
	//db.Where("code=?", query_name["code"]).First(&prod)
	//price, err := strconv.ParseUint((query_name["price"][0]), 10, 64)
	//db.Model(&prod).Update("Price", price)
	/* db.Model(&req).Update("Price", req.Price)
	db.Model(&req).Update("Code", req.Code)
	db.Model(&req).Update("Name", req.Name) */
	db.Model(&product.Product{}).Update("Price", req.Price)
	//fmt.Println(err)
	fmt.Println(db.Error)
}

func DeleteProduct(c *gin.Context) {
	prod := product.Product{}
	db := conn.GetPostgres()
	query_name := c.Request.URL.Query()
	db.Where("code=?", query_name["code"]).Delete(&prod)
	db.Unscoped().Where("code=?", query_name["code"]).Delete(&prod)
}
