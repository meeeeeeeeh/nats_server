package order_service

import (
	"nats_server/db"
	"nats_server/db/model"
)


const (
	filePath1 = "../model.json"
)


// вызов субскрайбера
// запись в кэш
// запись в бд
// гет запрос по айди заказа
// выдаем инфу по заказу



type OrderService struct {
	orderCash map[string]model.Order
	repository db.OrderRepository
	// subscriber
}

func InitOrderService() OrderService {
	db := db.NewOrderRepository()
	db
	return OrderService{
		repository: db,
		//new subscriber
	}
}

// вызываем методы репозитория



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