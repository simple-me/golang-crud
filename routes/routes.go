package routes
import (
       "net/http"
       products "CRUD-Operation/controllers/products"
       "github.com/gin-gonic/gin"
)
//StartGin function
func StartGin() {
       router := gin.Default()
       router.LoadHTMLGlob("templates/*")
       api := router.Group("/api")
       {
                api.POST("/product/create-product", products.CreateProduct)
                api.GET("/product/find-product/:code", products.FindProduct)
                api.GET("/product/list-products", products.ListProducts)
                api.POST("/product/update-product", products.UpdateProduct)
                api.POST("/product/delete-product", products.DeleteProduct)
       }
       router.NoRoute(func(c *gin.Context) {
              c.AbortWithStatus(http.StatusNotFound)
       })
       router.Run(":8000")
}
