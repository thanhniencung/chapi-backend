package model

type Cate struct {
	cateId  string `json:"cateId,omitempty" db:"cate_id,omitempty"`
	cateName string `json:"cateName,omitempty" db:"cate_name,omitempty"`
}
