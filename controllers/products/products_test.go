package user

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateProduct(t *testing.T) {
	arg := ProductParams{
		Name:  "santi20",
		Code:  "asdfh",
		Price: 23456,
	}
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(w)
	CreateProduct(c)

	fmt.Println(arg)
}
