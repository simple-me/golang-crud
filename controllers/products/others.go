package user

import (
	"CRUD-Operation/db/conn"
	product "CRUD-Operation/models/products"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HelloIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

type CreateAccountParams struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Price int64  `json:"price"`
}

func CreateProduct(c *gin.Context) {
	var req CreateAccountParams
	fmt.Println(c.Request.Body)
	/* arg := CreateAccountParams{
		Name:  req.Name,
		Code:  req.Code,
		Price: req.Price,
	} */
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	db := conn.GetPostgres()
	//db.AutoMigrate(&product.Product{})
	//query_name := c.Request.URL.Query()
	//fmt.Println(c.Request.URL.Query())
	/* price, err := strconv.ParseUint((query_name["price"][0]), 10, 64)
	if err != nil {
		fmt.Println("error in price value")
	} */
	//create := db.Create(&product.Product{Name: query_name["name"][0], Code: query_name["code"][0], Price: price})
	create := db.Create(&product.Product{Name: req.Name, Code: req.Code, Price: uint64(req.Price)})
	fmt.Println(create.Error)
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
	prod := product.Product{}
	db := conn.GetPostgres()
	query_name := c.Request.URL.Query()
	db.Where("code=?", query_name["code"]).First(&prod)
	price, err := strconv.ParseUint((query_name["price"][0]), 10, 64)
	db.Model(&prod).Update("Price", price)
	fmt.Println(err)
}

func DeleteProduct(c *gin.Context) {
	prod := product.Product{}
	db := conn.GetPostgres()
	query_name := c.Request.URL.Query()
	db.Where("code=?", query_name["code"]).Delete(&prod)
	db.Unscoped().Where("code=?", query_name["code"]).Delete(&prod)
}
