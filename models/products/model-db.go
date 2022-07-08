package model

import (
	"github.com/simple-me/golang-crud/db/conn"
)

var db = conn.GetPostgres()

func Get(code string) (*Product, error) {
	prod := Product{}
	err := db.First(&prod, "code=?", code).Error
	if err != nil {
		return nil, err
	}
	return &prod, nil
}

func GetAll() ([]Product, error) {
	prod := []Product{}
	err := db.Find(&prod).Error
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func Create(product Product) error {
	err := db.Create(&Product{Name: product.Name, Code: product.Code, Price: uint64(product.Price)}).Error
	if err != nil {
		return err
	}
	return nil
}

func Update(newProd Product) error {
	prod := Product{}
	err := db.First(&prod, "code=?", newProd.Code).Error
	if err != nil {
		return err
	}

	prod.Price = uint64(newProd.Price)
	prod.Name = newProd.Name

	db.Save(&prod)

	return nil
}

func Delete(code string) error {
	prod := Product{}
	err := db.First(&prod, "code=?", code).Error
	if err != nil {
		return err
	}

	err = db.Delete(&prod).Error
	if err != nil {
		return err
	}

	return nil
}
