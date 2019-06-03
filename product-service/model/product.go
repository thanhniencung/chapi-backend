package model

import "time"

type Product struct {
	ProductId  	 string    `json:"productId,omitempty" db:"product_id,omitempty""`
	ProductName  string    `json:"productName,omitempty" db:"product_name,omitempty" valid:"required"`
	ProductImage string    `json:"productImage,omitempty" db:"product_image,omitempty" valid:"required,url"`
	Quatity 	 int 	   `json:"quatity,omitempty" db:"quatity,omitempty" valid:"required,int"`
	SoldItems 	 int 	   `json:"soldItems,omitempty" db:"sold_items,omitempty"`
	Price 		 float64   `json:"price,omitempty" db:"price,omitempty" valid:"required,numeric"`
	CateId 		 string    `json:"cateId,omitempty" db:"cate_id,omitempty" valid:"required"`
	CreatedAt 	 time.Time `json:"createdAt,omitempty" db:"created_at,omitempty"`
	UpdatedAt	 time.Time `json:"updatedAt,omitempty" db:"updated_at,omitempty"`
	DeletedAt	 time.Time `json:"-"  db:"deleted_at,omitempty"`
}