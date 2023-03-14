package Repository

import (
	"proj-mido/stripe-gateway/Config"
	"proj-mido/stripe-gateway/Models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllProducts(products *[]Models.Products) (err error) {
	if err = Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *Models.Products) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}
