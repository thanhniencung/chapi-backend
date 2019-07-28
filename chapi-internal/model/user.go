package model

type User struct { // tag
	UserId      string `json:"user,omitempty" db:"user_id,omitempty"` // tags golang + parse json
	Role        string `json:"role,omitempty" db:"role,omitempty"`
	Phone       string `json:"phone,omitempty" db:"phone,omitempty" valid:"required"`
	Password    string `json:"password,omitempty" db:"password,omitempty" valid:"required"`
	DisplayName string `json:"displayName,omitempty" db:"display_name,omitempty" valid:"required"`
	Avatar      string `json:"avatar,omitempty" db:"avatar,omitempty" valid:"required"`
}