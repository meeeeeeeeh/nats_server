package order_service

import (
	"encoding/json"
	"fmt"
	"io"
	"nats_server/db"
	"nats_server/db/model"
	"os"
)

// вызов субскрайбера
// запись в кэш
// запись в бд
// гет запрос по айди заказа
// выдаем инфу по заказу

type OrderService struct {
	orderCash  map[string]model.Order
	repository db.OrderRepository
	// subscriber
}

func InitOrderService() OrderService {
	db, _ := db.NewOrderRepository()


	return OrderService{

		repository: db

		//new subscriber
	}
}




//func (o *OrderService) AddCashe()


func (o *OrderService) MsgProcess(m *stan.Msg){
	fmt.Println(m)

}



func GetFileData(filename string) (*model.Order, error) {
	var order model.Order

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() // выполнится  в люом случае либо в конце либо при панике (когда хз)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
