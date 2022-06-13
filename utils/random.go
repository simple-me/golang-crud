package utils

import (
	"math/rand"
	"time"
)

type ProductParams struct {
	Name  string
	Code  string
	Price int64
}

func RandomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomInt() int64 {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(300 - 100)
	return int64(v)
}

func RandomProductParams() ProductParams {
	prod := ProductParams{
		Name:  RandomString(4),
		Code:  RandomString(9),
		Price: RandomInt(),
	}

	return prod
}
