package entities

import (
	"fmt"
	"time"
)

type Product struct {
	Id       int    `gorm:"primary_key, AUTO_INCREAMENT"`
	Name     string `gorm: "column:name"`
	Price    float64
	Quanlity int
	Status   bool
	Created  time.Time
}

func (product *Product) TableName() string {
	return "product"
}

func (product Product) ToString() string {
	return fmt.Sprintf("\nId: %d\nname: %s\nPrice:%0.1f\nQuanlity: %d\nStatus: %t\nCreate: %s", product.Id, product.Name,product.Price, product.Quanlity, product.Status, product.Created.Format("02/01/2006"))
}

func (product ProductApi) ToString() string {
	return fmt.Sprintf("\nId: %d\nname: %s\nPrice:%0.1f", product.Id, product.Name,product.Price)
}
