package routes

import (
	"log"
	"net/http"

	products "github.com/simple-me/golang-crud/controllers/products"
	"github.com/simple-me/golang-crud/db/conn"

	"github.com/gin-gonic/gin"
)

//StartGin function
func StartGin() *gin.Engine {
	db := conn.GetPostgres()
	if db.Error != nil {
		log.Fatal(db.Error)
	}
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/product/create-product", products.CreateProduct)
		api.GET("/product/find-product/:code", products.FindProduct)
		api.GET("/product/list-products", products.ListProducts)
		api.POST("/product/update-product", products.UpdateProduct)
		api.DELETE("/product/delete-product", products.DeleteProduct)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	//router.Run(":8000")
	return router
}
