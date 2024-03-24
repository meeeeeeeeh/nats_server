package main

import (
<<<<<<< HEAD
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

	



=======
	"encoding/json"
	"fmt"
	"nats_server/config"
	"nats_server/db/model"

	"database/sql"
	"log"

	"io"
	"os"

	_ "github.com/lib/pq"


	"github.com/gobuffalo/pop"
	_ "github.com/lib/pq" // Драйвер базы данных PostgreSQL
)

//метод - функция которая принадлежит структуре
// у этой структуры есть методы по работе с бд
// 1 - по заполнению и 2 - поизвлечению
type OrderRepository struct {
	db *sql.DB
	//db connection
}


// закрытие - когда то в конце  // defer db.Close()
// нужно написать клозер - что будет если программа закончится но соединение с бд не закроется???



// метод репозитория 
func NewOrderRepository() (*OrderRepository, error) {
	db, err := sql.Open("postgres", config.DbConnect)
	if err != nil {
		return nil, err
    } 

	return &OrderRepository{
		db: db,
	}, nil
}


func (o *OrderRepository) AddOrder(order *model.OrderInfo) error {
	// insert into order_info
	statement := `INSERT INTO order_infoes (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := o.db.Exec(statement, order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	// insert to delivery
	statement = `INSERT INTO deliveries (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = o.db.Exec(statement, order.OrderUid, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return err
	}

	// insert into payment
	statement = `INSERT INTO payments (order_uid, trasaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = o.db.Exec(statement, order.OrderUid, order.Payment.Trasaction, order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err
	}

	// insert to items
	for idx, _ := range order.Items {
		statement = `INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
		_, err = o.db.Exec(statement, order.OrderUid, order.Items[idx].ChrtId, order.Items[idx].TrackNumber, order.Items[idx].Price, order.Items[idx].Rid, order.Items[idx].Name, order.Items[idx].Sale, order.Items[idx].Size, order.Items[idx].TotalPrice, order.Items[idx].NmId, order.Items[idx].Brand, order.Items[idx].Status)
		if err != nil {
			return err
		}
	}

	return nil
}


//метод для перемещения данных их б в кеш
func (o *OrderRepository) GetOrders() ([]model.OrderInfo, error) {
	//data := make(map[string]model.Order)

	var data []model.OrderInfo



	db, err := pop.Connect("development")
	if err != nil {
		return nil, err
	}

	log.Println("connection ok")

	err = db.All(&data)
	if err != nil {
		return nil, err
	}



	// // getting dat from order.info
	// rows, err := o.db.Query("SELECT * FROM order_info")
	// if err != nil {
	// 	// if err == sql.ErrNoRows {
	// 	// 	log.Fatalln("Db is empty")
	// 	// }
	// 	return nil, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order model.Order
	// 	err = rows.Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.Shardkey, &order.SmId, &order.DateCreated, &order.OofShard) 
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	log.Println(order)

	// 	data[order.OrderUid] = order
	// }




	// // getting dat from delivery
	// rows, err = o.db.Query("SELECT * FROM delivery")
	// if err != nil {
	// 	// if err == sql.ErrNoRows {
	// 	// 	log.Fatalln("Db is empty")
	// 	// }
	// 	return nil, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order model.Delivery
	// 	err = rows.Scan(order.Id, &order.OrderUid, &order.Name, &order.Phone, &order.Zip, &order.City, &order.Address, &order.Region, &order.Email) 
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	log.Println(order)

	// 	data[order.OrderUid] = order
	// }





	
	return data, nil
}



func GetFileData(filename string) (*model.OrderInfo, error) {
	var order model.OrderInfo

	file, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("db opening error")
		return nil, err
	}
	defer file.Close() // выполнится  в любом случае в конце и даже при панике 

	data, err := io.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("reading file error")
		return nil, err
	}

	err = json.Unmarshal(data, &order)
	if err != nil {
		err = fmt.Errorf("convertation json to struct")
		return nil, err
	}

	return &order, nil
>>>>>>> main
}



<<<<<<< HEAD
=======
func main() {

	orderRep, err := NewOrderRepository()
	if err != nil {
		log.Fatalln(err)
	}


	// adding order data to db

	// order, err := GetFileData(config.FilePath1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }


	// err = orderRep.AddOrder(order) 
	// if err != nil {
	// 	fmt.Println(err)
	// }




	// getting order data
	d, err := orderRep.GetOrders()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(d)
	


	

}


// Функция log.Fatal вызывается, когда программа встречает нечто непоправимое, такое как невозможность продолжения работы из-за ошибки. Она записывает сообщение в журнал и завершает программу.
// С другой стороны, функция log.Panic вызывается, когда программа столкнулась с ситуацией, которая не должна произойти, но потенциально может разрешиться. Она также записывает сообщение в журнал, но вместо завершения программы вызывает панику, что может быть обработано в коде с помощью функции recover().
//log.Fatalf("unable to connect to database: %v", err)
>>>>>>> main
//panic(err)
//recover
//Panic — это встроенная функция, которая останавливает обычный поток управления и начинает паниковать. Когда функция F вызывает panic, выполнение F останавливается, все отложенные вызовы в F выполняются нормально, затем F возвращает управление вызывающей функции. Для вызывающей функции вызов F ведёт себя как вызов panic. Процесс продолжается вверх по стеку, пока все функции в текущей го-процедуре не завершат выполнение, после чего аварийно останавливается программа. Паника может быть вызвана прямым вызовом panic, а также вследствие ошибок времени выполнения, таких как доступ вне границ массива.