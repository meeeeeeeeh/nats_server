package order_service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"nats_server/db"
	"nats_server/db/model"
)

type OrderService struct {
	repository db.OrderRepository
	orderCash  map[string]*model.Order
}

func InitOrderService() (*OrderService, error) {
	//db connection
	db, err := db.NewOrderRepository()
	if err != nil {
		return nil, err
	}

	//f there are some data in db -> add them to cash
	var cash map[string]*model.Order
	empty := db.IsEmpty()

	if !empty {
		cash, err = db.GetOrders()
		if err != nil {
			return nil, err
		}

	}

	return &OrderService{
		repository: *db,
		orderCash:  cash,
	}, nil
}

// new message processing
func MsgProcess(m *nats.Msg) {

	orderService, err := InitOrderService()
	if err != nil {
		log.Fatalln(err)
	}

	order, err := GetData(m.Data)
	if err != nil {
		log.Fatalln(err)
	}

	// adding data from this message to db
	orderService.repository.AddOrder(order)

	// updating cash with new data from db
	orderService.orderCash, err = orderService.repository.GetOrders()

}

func GetDataById(id string) (*model.Order, bool, error) {
	orderService, err := InitOrderService()
	if err != nil {
		return nil, false, err
	}

	// checks if there is smth in db
	// if there is no data in db the user gets message that there is no such order
	dbStatus := orderService.repository.IsEmpty()

	if dbStatus {
		return nil, false, nil
	}

	orderService.orderCash, err = orderService.repository.GetOrders()
	if err != nil {
		return nil, false, err
	}

	data, valid := orderService.orderCash[id]
	if !valid {
		return nil, false, nil
	}

	return data, true, nil
}

func GetData(data []byte) (*model.Order, error) {
	var order model.Order

	err := json.Unmarshal(data, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
