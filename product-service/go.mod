module chapi-backend/product-service

go 1.12

require (
	chapi-backend/chapi-internal v0.0.0-00010101000000-000000000000
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/labstack/echo v3.3.10+incompatible
)

replace chapi-backend/chapi-internal => ../chapi-internal
