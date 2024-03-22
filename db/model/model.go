package model

import (
	"time"
)

// то что с маленькой буквы видно в пределах только одного пакета
// с большой буквы видит вся программа

type Payment struct {
	// Id int
	Trasaction string `json:"transaction"`
	RequestId int `json:"request_id"`
	Currency string `json:"currency"`
	Provider string `json:"provider"`
	Amount int `json:"amount"`
	PaymentDt int `json:"payment_dt"`
	Bank string `json:"bank"`
	Delivery_cost int `json:"delivery_cost"`
	Goods_total int `json:"goods_total"`
	Custom_fee int `json:"custom_fee"`
}

type Item struct {
	// Id int
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


type Order struct {
	Id int
	OrderUid string `json:"order_uid"`
	TrackNumber int `json:"track_number"`
	Entry int `json:"entry"`
	Delivery Delivery `json:"delivery"`
	Payment Payment `json:"payment"`
	Item Item `json:"item"`
	// DeliveryId 
	// PaymentId int `json:"payment_id"`
	// ItemId int `json:"tem_id"`
	Locale int `json:"locale"`
	InternalSignature int `json:"internal_signature"`
	CustomerId int `json:"customer_id"`
	DeliveryService int `json:"delivery_service"`
	Shardkey int `json:"shardkey"`
	SmId int `json:"sm_id"`
	DateCreated time.Time `json:"date_created"`
	OofShard string `json:"oof_shard"`
}

type Delivery struct {
	// Id int 
	Name string `json:"page"`
	Phone string `json:"phone"`
	Zip string `json:"zip"`
	City string `json:"city"`
	Address string `json:"address"`
	Region string `json:"region"`
	Email string `json:"email"`
}