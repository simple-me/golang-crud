package user_test

import (
	"CRUD-Operation/routes"
	"CRUD-Operation/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

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
}

func TestCreateProduct(t *testing.T) {
	router := routes.StartGin()
	w := httptest.NewRecorder()

	var buf bytes.Buffer
	prod := utils.RandomProductParams()
	err := json.NewEncoder(&buf).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "/api/product/create-product", &buf)
	router.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
}

func TestFindProduct(t *testing.T) {
	router := routes.StartGin()
	w := httptest.NewRecorder()

	//Create random product
	var buf bytes.Buffer
	prod := utils.RandomProductParams()
	err := json.NewEncoder(&buf).Encode(prod)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/product/create-product", &buf)
	require.NoError(t, err)
	router.ServeHTTP(w, req)

	//Find that same recently created random product
	req, err = http.NewRequest("GET", fmt.Sprintf("/api/product/find-product/%s", prod.Code), nil)
	require.NoError(t, err)
	router.ServeHTTP(w, req)

	data, err := ioutil.ReadAll(w.Body)
	require.NoError(t, err)

	var r ProductParams
	json.Unmarshal(data, &r)
	fmt.Println(r.Response.Code)
	require.Equal(t, r.Response.Code, prod.Code)
}
