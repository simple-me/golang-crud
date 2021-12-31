package model

import (
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model `json:"-"`
  Name string `gorm:"unique";not null`
  Code  string `gorm:"unique";not null`
  Price uint64
}
