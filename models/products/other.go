package model

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Price uint64 `json:"price"`
}
