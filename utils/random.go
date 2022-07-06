package utils

import (
	"math/rand"
	"time"
)

type ProductParams struct {
	Name  string
	Code  string
	Price interface{}
}

func RandomString(str string) string {

	rand.Seed(time.Now().Unix())

	//str := "thisisarandomstring123"

	shuff := []rune(str)

	// Shuffling the string
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	// Displaying the random string
	return string(shuff)
}

func RandomInt() int64 {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(300 - 100)
	return int64(v)
}

func RandomProductParams(s string) ProductParams {
	prod := ProductParams{
		Name:  RandomString(s),
		Code:  RandomString(s),
		Price: RandomInt(),
	}

	return prod
}
