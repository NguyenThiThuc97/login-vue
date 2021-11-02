package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primary_key"`
	ProductName string `json:"productname" gorm:not null`
	DelFlg      uint   `json:"del_fg" gorm:"default:1"`
}

type CreateProduct struct {
	ProductName string `json:"productname" gorm:not null`
}

func (product *Product) Create() (*Product, error) {
	err := DB.Create(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}

func DeleteProduct(id uint) error {
	product := Product{}
	result := DB.Model(&Product{}).Where("id = ?", id).Find(&product)
	err := result.Error
	if err != nil || result.RowsAffected == 0 {
		return err
	}
	DB.Model(&Product{}).Where("id = ?", id).Update("del_flg", 0)
	return err
}
