package models

import (
	"config"
	"entities"
)

type ProductModel struct {
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) Condition1(status bool) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Where("status = ? ", status).Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) Condition2(min, max float64) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Where("price >= ? and price <= ? ", min, max).Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) Like(keyword string) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Where("name like ?", "%"+keyword+"%").Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) Sort() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Order("price desc").Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) SortByCondition(status bool) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Where("status =? ", status).Order("price desc").Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) Limit(n int) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Limit(n).Order("price desc").Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) LimitOrder(n int) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Order("id desc").Limit(n).Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) LimitOrderCondition(status bool, n int) ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		db.Where("status = ?", status).Order("id desc").Limit(n).Find(&products)
		return products, nil
	}
}

func (productModel ProductModel) FindById(id int) (entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return entities.Product{}, err
	} else {
		var product entities.Product
		db.Where("id = ?", id).First(&product)
		return product, nil
	}
}

func (productModel ProductModel) FindAllApi() ([]entities.ProductApi, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		var productsApi []entities.ProductApi
		db.Table("product").Select("id, name, price").Scan(&productsApi)
		return productsApi, nil
	}
}

func (producModel ProductModel) Count() int64 {
	db, _ := config.GetDB()
	var result int64
	db.Table("product").Count(&result)
	return result
}

func (producModel ProductModel) CountByStatus(status bool) int64 {
	db, _ := config.GetDB()
	var result int64
	db.Table("product").Where("status = ?", status).Count(&result)
	return result
}

func (producModel ProductModel) SumProduct() int64 {
	db, _ := config.GetDB()
	var total int64
	rows, _ := db.Table("product").Select("(quantity) as total").Rows()
	rows.Next()
	rows.Scan(&total)
	return total
}

func (producModel ProductModel) MaxAndMinPrice() (float64, float64) {
	db, _ := config.GetDB()
	var min, max float64
	rows, _ := db.Table("product").Select("min(price) as minPrice, max(price) as maxPrice").Rows()
	rows.Next()
	rows.Scan(&min, &max)
	return min, max
}


func (productModel ProductModel) CreateProduct(product *entities.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		db.Create(&product)
		return nil
	}
}

func (productModel ProductModel) UpdateProduct(product *entities.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		db.Save(&product)
		return nil
	}
}

func (productModel ProductModel) DeleteProduct(product entities.Product) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	} else {
		db.Delete(&product)
		return nil
	}
}