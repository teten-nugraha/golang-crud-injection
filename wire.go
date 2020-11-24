package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/teten-nugraha/golang-crud-injection/product"
)

func initProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProviderProductRepository, product.ProviderProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}
