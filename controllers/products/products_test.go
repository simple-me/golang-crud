package user_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/simple-me/golang-crud/routes"
	"github.com/simple-me/golang-crud/utils"

	"github.com/stretchr/testify/require"
)

type ProductParams struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Price    int64  `json:"price"`
	Response Resp
}

type Resp struct {
	Name  string
	Code  string
	Price int64
}

func TestListProducts(t *testing.T) {
	router := routes.StartGin()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/product/list-products", nil)
	router.ServeHTTP(w, req)

	fmt.Println(w.Body)
	require.Equal(t, 200, w.Code)

	//
	w2 := httptest.NewRecorder()
	os.Setenv("PG_CONNSTRING", "postgres://root:secret@127.0.0.1:1000/products")
	fmt.Println(os.Getenv("PG_CONNSTRING"))
	req, _ = http.NewRequest("GET", "/api/product/list-products", nil)
	router.ServeHTTP(w2, req)

	fmt.Println(w2.Body)
	require.Equal(t, 500, w2.Code)
	os.Unsetenv("PG_CONNSTRING")
}

func TestCreateProduct(t *testing.T) {
	router := routes.StartGin()

	//Create random product - valid values
	w := httptest.NewRecorder()
	var buf bytes.Buffer
	prod := utils.RandomProductParams("thisisarandomstring123456")
	err := json.NewEncoder(&buf).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "/api/product/create-product", &buf)
	router.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)

	//Create random product - invalid value in json
	w2 := httptest.NewRecorder()
	var buf2 bytes.Buffer
	prod2 := utils.RandomProductParams("thisisarandomstring123")
	prod2.Price = "invalid value"
	err = json.NewEncoder(&buf2).Encode(prod2)
	if err != nil {
		log.Fatal(err)
	}
	req, _ = http.NewRequest("POST", "/api/product/create-product", &buf2)
	router.ServeHTTP(w2, req)

	require.Equal(t, 400, w2.Code)

	//Create random product - repeated value, unique constraints, pick values from first case
	w3 := httptest.NewRecorder()
	var buf3 bytes.Buffer
	err = json.NewEncoder(&buf3).Encode(prod)
	req, _ = http.NewRequest("POST", "/api/product/create-product", &buf3)
	router.ServeHTTP(w3, req)
	require.Equal(t, 500, w3.Code)
}

func TestFindProduct(t *testing.T) {
	router := routes.StartGin()
	w := httptest.NewRecorder()

	//Create random product
	var buf bytes.Buffer
	prod := utils.RandomProductParams("thisisarandomstring123456789")
	err := json.NewEncoder(&buf).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/product/create-product", &buf)
	require.NoError(t, err)
	router.ServeHTTP(w, req)

	//Find that same recently created random product -- product found
	req, err = http.NewRequest("GET", fmt.Sprintf("/api/product/find-product/%s", prod.Code), nil)
	require.NoError(t, err)
	router.ServeHTTP(w, req)

	data, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	var r ProductParams
	json.Unmarshal(data, &r)
	fmt.Println(r.Response.Code)
	require.Equal(t, prod.Code, r.Response.Code)

	//Find that same recently created random product -- product not found
	req, err = http.NewRequest("GET", "/api/product/find-product/blah", nil)
	router.ServeHTTP(w, req)

	require.NoError(t, err)
	msg, err := strconv.Unquote(w.Body.String())
	require.NoError(t, err)
	require.Equal(t, "record not found", msg)

}

func TestDeleteProduct(t *testing.T) {
	router := routes.StartGin()

	//Create random product
	w := httptest.NewRecorder()
	var buf bytes.Buffer
	prod := utils.RandomProductParams("thisisarandomstring1234567890123")
	err := json.NewEncoder(&buf).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/product/create-product", &buf)
	require.NoError(t, err)
	router.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	//Delete recently created product
	var buf2 bytes.Buffer
	w2 := httptest.NewRecorder()
	err = json.NewEncoder(&buf2).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, err = http.NewRequest("DELETE", "/api/product/delete-product", &buf2)
	require.NoError(t, err)
	router.ServeHTTP(w2, req)
	require.Equal(t, 200, w2.Code)

	// Detete product - invalid value - not found
	var buf3 bytes.Buffer
	w3 := httptest.NewRecorder()
	err = json.NewEncoder(&buf3).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, err = http.NewRequest("DELETE", "/api/product/delete-product", &buf3)
	require.NoError(t, err)
	router.ServeHTTP(w3, req)
	require.Equal(t, 404, w3.Code)
}

func TestUpdateProduct(t *testing.T) {
	router := routes.StartGin()
	w := httptest.NewRecorder()

	//Create random product
	var buf bytes.Buffer
	prod := utils.RandomProductParams("thisisarandomstring156")
	err := json.NewEncoder(&buf).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/product/create-product", &buf)
	require.NoError(t, err)
	router.ServeHTTP(w, req)

	//Update recently created product
	prod.Name = utils.RandomString("thisisarandomstring156")
	prod.Price = utils.RandomInt()
	err = json.NewEncoder(&buf).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, err = http.NewRequest("POST", "/api/product/update-product", &buf)
	require.NoError(t, err)
	router.ServeHTTP(w, req)
}
