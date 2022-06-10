package model

type Product struct {
	//gorm.Model `json:"-"`
	ID    uint   `gorm:"primarykey"`
	Name  string `gorm:"unique";not null`
	Code  string `gorm:"unique";not null`
	Price uint64
}
