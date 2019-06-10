package repo_impl

import (
	"chapi-backend/chapi-internal/db"
	"chapi-backend/order-service/model"
	"chapi-backend/order-service/repository"
	internal "chapi-backend/chapi-internal/encrypt"
	"context"
	"database/sql"
	"errors"
	"time"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) repository.OrderRepository {
	return &OrderRepoImpl{
		sql: sql,
	}
}

// 1. Insert 1 record vào order table
// chú ý trước khi tạo 1 record trong order table thì kiểm tra user hiện tại đã có
// order hay chưa, nếu có rồi thì ko làm gì cả, chưa có thì mới tạo mới
// 2. insert 1 record vào card table
func (o *OrderRepoImpl) AddToCard(context context.Context, userId string, card model.Card) error {
	sqlCheckOrder := `select exists(select 1 from orders where user_id = $1 and status = $2)`
	var orderRow = model.Order{}
	err := o.sql.Db.GetContext(context, &orderRow, sqlCheckOrder, userId, model.ORDERING.String())
	if err != nil && err == sql.ErrNoRows {
		// Tạo 1 order mới mới với staus = ORDERING
		sqlInsertOrderStatement := `
		  INSERT INTO orders(user_id, order_id, status, updated_at) 
          VALUES(:user_id, :order_id, :status, :updated_at)
     	`
		orderRow.UserId = userId
		orderRow.OrderId = internal.UUID()
		orderRow.UpdatedAt = time.Now()
		orderRow.Status = model.ORDERING.String()

		_, err := o.sql.Db.NamedExecContext(context, sqlInsertOrderStatement, orderRow)
		if err != nil {
			return err
		}
	}

	// Nếu đã tạo order cho user này rồi thì chỉ update table card thôi
	sqlInsertCardStatement := `
		  INSERT INTO card(order_id, product_id, product_name, product_image, quantity, price) 
          VALUES(:order_id, :product_id, :product_name, :product_image, :quantity, :price)
     `
	card.OrderId = orderRow.OrderId
	_, err = o.sql.Db.NamedExecContext(context, sqlInsertCardStatement, card)

	return err
}

func (o *OrderRepoImpl) UpdateStateOrder(context context.Context, order model.Order) error {
	// status, order_id, user_id
	sqlStatement := `
		UPDATE orders
		SET 
			status = :status
			updated_at = :updated_at
		WHERE 
			user_id = :user_id 
			AND order_id = :order_id
	`

	order.UpdatedAt = time.Now()
	result, err := o.sql.Db.NamedExecContext(context, sqlStatement, order)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Update thất bại")
	}

	return nil
}

func (o *OrderRepoImpl) CountShoppingCard(context context.Context, userId string) (model.OrderCount, error) {
	// Order của ai thì người đó mới được xem thông tin : orders.user_id = $2
	sqlCountStatement := `
		SELECT
		   orders.order_id,	
		   COUNT(*) AS count_item
		FROM
		   orders
		INNER JOIN card 
		ON 
		  orders.user_id = $1 AND 
		  orders.order_id = card.order_id AND 
		  orders.status = 'ORDERING'
		GROUP BY
		   orders.order_id
	`
	row := o.sql.Db.QueryRowxContext(context, sqlCountStatement, userId)

	orderCount := model.OrderCount{}

	err := row.Err()
	if err != nil {
		return orderCount, err
	}

	err = row.StructScan(&orderCount)
	if err != nil {
		return orderCount, err
	}

	return orderCount, nil
}

func (o *OrderRepoImpl) ShoppingCard(context context.Context, userId string, orderId string) (model.Order, error) {
	sqlShoppingCard := `
		SELECT
		   orders.order_id,
		   card.product_id,
		   card.product_name,
		   card.product_image,
		   card.quantity,
		   card.price
		FROM
		   orders
		INNER JOIN card 
		ON 
          orders.user_id = $1 AND 
		  orders.order_id = $2 AND 
		  orders.order_id = card.order_id AND 
		  orders.status = 'ORDERING'
	`
	orders := model.Order{}
	cards := []model.Card{}

	err := o.sql.Db.SelectContext(context, &cards, sqlShoppingCard, userId, orderId)
	if err != nil {
		return orders, err
	}

	var sum float64 = 0
	for _, card := range cards {
		sum += card.Price * float64(card.Quantity)
	}

	orders.Total = sum
	orders.Items = cards

	return orders, nil
}

