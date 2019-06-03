module chapi-backend/user-service

go 1.12

require (
	chapi-backend/chapi-internal v1.0.0
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a

	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.8
	github.com/lib/pq v1.1.1
	github.com/stretchr/testify v1.3.0 // indirect
)

replace chapi-backend/chapi-internal => ../chapi-internal
