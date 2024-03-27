package db

import (
	"nats_server/config"
	"nats_server/db/model"

	"database/sql"
	_ "github.com/lib/pq"
)

type OrderRepository struct {
	db *sql.DB
	//db connection
}

func NewOrderRepository() (*OrderRepository, error) {
	db, err := sql.Open("postgres", config.DbConnect)
	if err != nil {
		return nil, err
	}

	return &OrderRepository{
		db: db,
	}, nil
}

func (o *OrderRepository) IsEmpty() bool {

	rows, err := o.db.Query("SELECT * FROM order_info")
	if err != nil || !rows.Next() {
		if err == sql.ErrNoRows {
			return true
		}
		return true
	}
	defer rows.Close()
	return false

}

func (o *OrderRepository) AddOrder(order *model.Order) error {
	// insert into order_info
	statement := `INSERT INTO order_info (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := o.db.Exec(statement, order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	// insert to delivery
	statement = `INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = o.db.Exec(statement, order.OrderUid, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return err
	}

	// insert into payment
	statement = `INSERT INTO payment (order_uid, trasaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = o.db.Exec(statement, order.OrderUid, order.Payment.Trasaction, order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err
	}

	// insert to items
	for idx, _ := range order.Items {
		statement = `INSERT INTO item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
		_, err = o.db.Exec(statement, order.OrderUid, order.Items[idx].ChrtId, order.Items[idx].TrackNumber, order.Items[idx].Price, order.Items[idx].Rid, order.Items[idx].Name, order.Items[idx].Sale, order.Items[idx].Size, order.Items[idx].TotalPrice, order.Items[idx].NmId, order.Items[idx].Brand, order.Items[idx].Status)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *OrderRepository) GetOrders() (map[string]*model.Order, error) {
	data := make(map[string]*model.Order)

	// getting data from order.info
	rows, err := o.db.Query("SELECT * FROM order_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.Shardkey, &order.SmId, &order.DateCreated, &order.OofShard)
		if err != nil {
			return nil, err
		}
		data[order.OrderUid] = &order
	}

	// getting data from delivery
	rows, err = o.db.Query("SELECT * FROM delivery")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order model.Delivery
		err = rows.Scan(&order.Id, &order.OrderUid, &order.Name, &order.Phone, &order.Zip, &order.City, &order.Address, &order.Region, &order.Email)
		if err != nil {
			return nil, err
		}
		data[order.OrderUid].Delivery = order
	}

	// getting data from payment
	rows, err = o.db.Query("SELECT * FROM payment")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order model.Payment
		err = rows.Scan(&order.Id, &order.OrderUid, &order.Trasaction, &order.RequestId, &order.Currency, &order.Provider, &order.Amount, &order.PaymentDt, &order.Bank, &order.DeliveryCost, &order.GoodsTotal, &order.CustomFee)
		if err != nil {
			return nil, err
		}
		data[order.OrderUid].Payment = order
	}

	// getting data from items
	rows, err = o.db.Query("SELECT * FROM item")
	if err != nil {
		return nil, err
	}

	var items []model.Item
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.Id, &item.OrderUid, &item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if len(items) == 1 {
		data[items[0].OrderUid].Items = items

	} else {
		var forEachOrder []model.Item

		for idx, _ := range items {
			i := 0

			if idx == 0 || items[idx].OrderUid == items[idx-1].OrderUid {
				forEachOrder = append(forEachOrder, items[idx])
				i++
			} else {
				data[items[i].OrderUid].Items = forEachOrder
				forEachOrder = nil
				forEachOrder = append(forEachOrder, items[idx])
				i = 0
			}
		}
		data[items[len(items)-1].OrderUid].Items = forEachOrder
	}

	return data, nil
}
