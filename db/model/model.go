package model

import (
	"time"
)

// то что с маленькой буквы видно в пределах только одного пакета
// с большой буквы видит вся программа

type Payment struct {
	Id int
	OrderId string
	Trasaction string `json:"transaction"`
	RequestId string `json:"request_id"`
	Currency string `json:"currency"`
	Provider string `json:"provider"`
	Amount int `json:"amount"`
	PaymentDt int `json:"payment_dt"`
	Bank string `json:"bank"`
	DeliveryCost int `json:"delivery_cost"`
	GoodsTotal int `json:"goods_total"`
	CustomFee int `json:"custom_fee"`
}

type Item struct {
	Id int 
	OrderId string
	ChrtId int `json:"chrt_id"`
    TrackNumber string `json:"track_number"`
    Price int `json:"price"`
    Rid string `json:"rid"`
    Name string `json:"name"`
    Sale int `json:"sale"`
    Size string `json:"size"`
    TotalPrice int `json:"total_price"`
    NmId int `json:"nm_id"`
    Brand string `json:"brand"`
    Status int `json:"status"`
}

type Delivery struct {
	Id int 
	OrderUid string
	Name string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
	Zip string `json:"zip" db:"zip"`
	City string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region string `json:"region" db:"region"`
	Email string `json:"email" db:"email"`
}

type OrderInfo struct {
	OrderUid string `json:"order_uid" db:"order_uid"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Entry string `json:"entry" db:"entry"`
	Delivery Delivery `json:"delivery"`
	Payment Payment `json:"payment"`
	Items []Item `json:"items"`
	Locale string `json:"locale" db:"locale"`
	InternalSignature string `json:"internal_signature" db:"internal_signature"`
	CustomerId string `json:"customer_id" db:"customer_id"`
	DeliveryService string `json:"delivery_service" db:"delivery_service"`
	Shardkey string `json:"shardkey" db:"shardkey"`
	SmId int `json:"sm_id" db:"sm_id"`
	DateCreated time.Time `json:"date_created" db:"date_created"`
	OofShard string `json:"oof_shard" db:"oof_shard"`
}
