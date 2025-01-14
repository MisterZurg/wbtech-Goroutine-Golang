package repository

import (
	"github.com/MisterZurg/wbtech-Goroutine-Golang/l0/internal/handlers"
	"github.com/jmoiron/sqlx"
)

type OrdersRepository struct {
	*sqlx.DB
}

func (or *OrdersRepository) CreateOrder(order handlers.Order) {
	// tx, err := or.DB.Begin()
	queryOrders := `INSERT INTO orders
            (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	rowOrder := or.DB.QueryRow(
		queryOrders,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.ShardKey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)

	var id int64
	if err := rowOrder.Scan(&id); err != nil {
		// TODO cancel tx
	}

	queryDeliveries := `INSERT INTO deliveries (order_uid, name, phone, zip, city, address, region, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	rowDeliveries := or.DB.QueryRow(
		queryDeliveries,
		order.OrderUID,
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email,
	)

	if err := rowDeliveries.Scan(&id); err != nil {
		// TODO cancel tx
	}

	queryItems := `INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, i_size, total_price, nm_id, brand, status)`

	for _, item := range order.Items {
		// TODO: improvement batch/bulk insert
		rowItem := or.DB.QueryRow(
			queryItems,
			order.OrderUID,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status,
		)
		if err := rowItem.Scan(&id); err != nil {
			// TODO cancel tx
		}
	}

	queryPayments := `INSERT INTO payments (transaction_id, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	rowItem := or.DB.QueryRow(
		queryPayments,
		order.OrderUID,
		order.Payment.RequestID,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.PaymentDT,
		order.Payment.Bank,
		order.Payment.DeliveryCost,
		order.Payment.GoodsTotal,
		order.Payment.CustomFee,
	)
	if err := rowItem.Scan(&id); err != nil {
		// TODO cancel tx
	}

	return
}

func (or *OrdersRepository) GetOrder(orderUID string) handlers.Order {
	order := &handlers.Order{}

	queryOrders := "SELECT * FROM orders WHERE order_uid = $1"
	// err = or.DB.QueryRowx(`SELECT courier_id, courier_type, array_to_json(regions), array_to_json(working_hours) FROM couriers WHERE courier_id=$1`, courierId).Scan(
	err := or.DB.QueryRowx(queryOrders, orderUID).
		Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
		)
	if err != nil {
		// TODO: return internal err
	}
	//json.Unmarshal([]byte(regsStr), &ans.Regions)
	//json.Unmarshal([]byte(wohoStr), &ans.WorkingHours)

	queryDelivery := `SELECT * FROM deliveries WHERE order_uid = $1`
	err = or.DB.QueryRowx(queryDelivery, orderUID).
		Scan(
			&order.Delivery.Name,
			&order.Delivery.Phone,
			&order.Delivery.Zip,
			&order.Delivery.City,
			&order.Delivery.Address,
			&order.Delivery.Region,
			&order.Delivery.Email,
		)
	if err != nil {
		// TODO: return internal err
	}

	queryPaymentment := `SELECT * FROM payments WHERE transaction_id = $1`
	err = or.DB.QueryRowx(queryPaymentment, orderUID).
		Scan(
			&order.Payment.Transaction,
			&order.Payment.RequestID,
			&order.Payment.Currency,
			&order.Payment.Provider,
			&order.Payment.Amount,
			&order.Payment.PaymentDT,
			&order.Payment.Bank,
			&order.Payment.DeliveryCost,
			&order.Payment.GoodsTotal,
			&order.Payment.CustomFee,
		)
	if err != nil {
		// TODO: return internal err
	}

	items := make([]handlers.Item, 0)
	queryItems := `SELECT * FROM items WHERE order_uid = $1`
	rows, err := or.DB.Queryx(queryItems, orderUID)
	for rows.Next() {
		var item handlers.Item
		err = rows.StructScan(&item)
		if err != nil {
			// TODO: return internal err
		}
		items = append(items, item)
	}
	order.Items = items
	return *order
}
