module chapi-backend/order-service

go 1.12

require (
	chapi-backend/chapi-internal v1.0.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/stretchr/testify v1.3.0 // indirect
	google.golang.org/appengine v1.6.1 // indirect
)

replace chapi-backend/chapi-internal => ../chapi-internal
