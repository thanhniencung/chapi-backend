module chapi-backend/user-service

go 1.12

require (
	chapi-backend/chapi-internal v1.0.0
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a

	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.8
	github.com/lib/pq v1.1.1
	github.com/nyaruka/phonenumbers v1.0.42
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/ttacon/builder v0.0.0-20170518171403-c099f663e1c2 // indirect
	github.com/ttacon/libphonenumber v1.0.1
)

replace chapi-backend/chapi-internal => ../chapi-internal
