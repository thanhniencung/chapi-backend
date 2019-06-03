package model

type User struct {
	UserId 		string `json:"userId,omitempty" db:"userId,omitempty"`
	Phone 		string `json:"phone,omitempty" db:"phone,omitempty" valid:"required"`
	Password 	string `json:"password,omitempty" db:"password,omitempty" valid:"required"`
	DisplayName string `json:"displayName,omitempty" db:"displayName,omitempty" valid:"required"`
	Avatar 		string `json:"avatar,omitempty" db:"avatar,omitempty" valid:"required"`
}
