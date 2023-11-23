package usecase

import (
	"errors"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/repository"
)

// Category Management----------------------------------------------------------------------------

func AddCategories(categoryData *domain.Category) error {
	err := repository.AddCategories(categoryData)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(updateCategory domain.Category, id int) error {
	err := repository.UpdateCategory(updateCategory, id)
	if err != nil {
		return err
	}
	return nil

}

func DeleteCategory(id int) error {
	err := repository.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

// Brand------------------------------------------------------------------------------------------

func AddBrand(newBrand domain.Brand) error {
	err := repository.AddBrand(newBrand)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBrand(updatedBrand domain.Brand, id int) error {
	err := repository.UpdateBrand(updatedBrand, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBrand(id int) error {
	if err := repository.DeleteBrand(id); err != nil {
		return err
	}
	return nil
}

// Product Management------------------------------------------------------------------------------

func AddProduct(productData *domain.Product) error {
	res, err := repository.FindProductByProductName(productData)
	if err != nil && res == nil {
		err = repository.AddProduct(productData)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Product already existed")
}

func UpdateProduct(updateProduct domain.Product, id int) error {
	err := repository.UpdateProduct(updateProduct, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int) error {
	if err := repository.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}

// User- Products---------------------------------------------------------------------

func GetProducts(limit, offset int) ([]domain.Product, error) {
	products, err := repository.GetProducts(limit, offset)
	if err != nil {
		return nil, err
	}
	return products, err
}

func GetProductByID(id int) (*domain.Product, error) {
	product, err := repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductByBrandID(limit, offset, id int) ([]domain.Product, error) {
	product, err := repository.GetProductByBrandID(limit, offset, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductByName(name string) (*domain.Product, error) {
	product, err := repository.GetProductByName(name)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductCategoryID(limit, offset, id int) ([]domain.Product, error) {
	product, err := repository.GetProductByCategoryID(limit, offset, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
