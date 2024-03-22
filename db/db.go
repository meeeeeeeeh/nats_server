package main

import (
	//"time"
	"encoding/json"
	"os"
	"github.com/meeeeeeeeh/nats_server/db/model"
	"fmt"
	"io"
	//"log"
)

// type Payment struct {
// 	// Id int
// 	Trasaction string `json:"transaction"`
// 	RequestId string `json:"request_id"`
// 	Currency string `json:"currency"`
// 	Provider string `json:"provider"`
// 	Amount int `json:"amount"`
// 	PaymentDt int `json:"payment_dt"`
// 	Bank string `json:"bank"`
// 	Delivery_cost int `json:"delivery_cost"`
// 	Goods_total int `json:"goods_total"`
// 	Custom_fee int `json:"custom_fee"`
// }

// type Item struct {
// 	// Id int
// 	ChrtId int `json:"chrt_id"`
//     TrackNumber string `json:"track_number"`
//     Price int `json:"price"`
//     Rid string `json:"rid"`
//     Name string `json:"name"`
//     Sale int `json:"sale"`
//     Size string `json:"size"`
//     TotalPrice int `json:"total_price"`
//     NmId int `json:"nm_id"`
//     Brand string `json:"brand"`
//     Status int `json:"status"`
// }


// type Order struct {
// 	// Id int
// 	OrderUid string `json:"order_uid"`
// 	TrackNumber string `json:"track_number"`
// 	Entry string `json:"entry"`
// 	Delivery Delivery `json:"delivery"`
// 	Payment Payment `json:"payment"`
// 	Item Item `json:"item"`
// 	// DeliveryId 
// 	// PaymentId int `json:"payment_id"`
// 	// ItemId int `json:"tem_id"`
// 	Locale string `json:"locale"`
// 	InternalSignature string `json:"internal_signature"`
// 	CustomerId string `json:"customer_id"`
// 	DeliveryService string `json:"delivery_service"`
// 	Shardkey string `json:"shardkey"`
// 	SmId int `json:"sm_id"`
// 	DateCreated time.Time `json:"date_created"`
// 	OofShard string `json:"oof_shard"`
// }

// type Delivery struct {
// 	// Id int 
// 	Name string `json:"page"`
// 	Phone string `json:"phone"`
// 	Zip string `json:"zip"`
// 	City string `json:"city"`
// 	Address string `json:"address"`
// 	Region string `json:"region"`
// 	Email string `json:"email"`
// }

func main() {
	var srt Order
	f, err := os.Open("../model.json")
	if err != nil {
		panic(err)
	}

	defer f.Close() // выполнится  в люом случае либо в конце либо при панике (когда хз)

	file, err := io.ReadAll(f)
	//fmt.Println(string(file))

	err = json.Unmarshal(file, &srt)

	fmt.Println(srt.CustomerId)

	



}



//panic(err)
//recover
//Panic — это встроенная функция, которая останавливает обычный поток управления и начинает паниковать. Когда функция F вызывает panic, выполнение F останавливается, все отложенные вызовы в F выполняются нормально, затем F возвращает управление вызывающей функции. Для вызывающей функции вызов F ведёт себя как вызов panic. Процесс продолжается вверх по стеку, пока все функции в текущей го-процедуре не завершат выполнение, после чего аварийно останавливается программа. Паника может быть вызвана прямым вызовом panic, а также вследствие ошибок времени выполнения, таких как доступ вне границ массива.