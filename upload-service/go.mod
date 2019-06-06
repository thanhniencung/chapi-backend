module upload-service

go 1.12

require (
	chapi-backend/chapi-internal v1.0.0
	github.com/labstack/echo v3.3.10+incompatible
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5 // indirect
)

replace chapi-backend/chapi-internal => ../chapi-internal
