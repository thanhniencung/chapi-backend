package handler

import (
	"chapi-backend/order-service/repository"
	"github.com/labstack/echo"
)

type OrderHandler struct {
	OrderRepo repository.OrderRepository
}

// Click nút mua là call api này
func (m *OrderHandler) AddToCard(c echo.Context) error {
	return nil
}

// Click vào shopping card icon ở AppBar sẽ call api này
func (m *OrderHandler) OrderDetails(c echo.Context) error {
	return nil
}

// Xác nhận order này đã chuẩn ở phía user
func (m *OrderHandler) ConfirmOrder(c echo.Context) error {
	return nil
}

// Khi vào app gọi api này để hiển thị số lượng sản phẩm có trong shopping card
// Hiển thị ở phần icon AppBar
func (m *OrderHandler) OrderCountItem(c echo.Context) error {
	return nil
}
