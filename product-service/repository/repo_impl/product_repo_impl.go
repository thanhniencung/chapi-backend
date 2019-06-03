package repo_impl

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/product-service/model"
	"chapi-backend/product-service/repository"
	"context"
	"time"
)

type ProductRepoImpl struct {
	sql *db.Sql
}

func NewProductRepo(sql *db.Sql) repository.ProductRepository {
	return &ProductRepoImpl{
		sql: sql,
	}
}

func (u *ProductRepoImpl) AddProduct(context context.Context, product model.Product) (model.Product, error) {
	sqlStatement := `
		  INSERT INTO product(
		  		product_id, product_name, product_image, quatity, 
		  		sold_items, created_at, updated_at, price, cate_id) 
          VALUES(:product_id, :product_name, :product_image, :quatity, 
          		 :sold_items, :created_at, :updated_at, :price, :cate_id)
     `

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, sqlStatement, product)
	return product, err
}

func (u *ProductRepoImpl) UpdateProduct(context context.Context, product model.Product) (model.Product, error) {
	sqlStatement := `
		UPDATE product
		SET product_name = :product_name, 
			quatity = :quatity, 
			sold_items = :sold_items, 
			price = :price, 
			cate_id = :cate_id, 
			updated_at = :updated_at
		WHERE product_id = :product_id;
	`

	product.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, sqlStatement, product)
	return product, err
}

func (u *ProductRepoImpl) DeletePRoduct(context context.Context, productId string) (error) {
	sqlStatement := ` 
		UPDATE product
		SET deleted_at = $1
		WHERE product_id = $2;
	`
	_, err := u.sql.Db.ExecContext(context, sqlStatement, time.Now(), productId)
	return err
}

func (u *ProductRepoImpl) SelectProductById(context context.Context, productId string) (model.Product, error) {
	var product model.Product

	row := u.sql.Db.QueryRowxContext(context, "SELECT * FROM product WHERE product_id=$1", productId)

	err := row.Err()
	if err != nil {
		return product, err
	}

	err = row.StructScan(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}



