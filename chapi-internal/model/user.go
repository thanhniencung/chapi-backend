package model

type User struct {
	UserId 		string `json:"userId,omitempty" db:"user_id,omitempty"`
	Phone 		string `json:"phone,omitempty" db:"phone,omitempty"`
	Password 	string `json:"password,omitempty" db:"password,omitempty"`
	Role 		string `json:"role,omitempty" db:"role,omitempty"`
	DisplayName string `json:"displayName,omitempty" db:"display_name,omitempty"`
	Avatar 		string `json:"avatar,omitempty" db:"avatar,omitempty"`
}
