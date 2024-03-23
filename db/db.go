package main

import (
	//"time"
	"encoding/json"
	"fmt"
	"nats_server/db/model"
	"os"

	//"fmt"
	"database/sql"
	"io"
	"log"

	_ "github.com/lib/pq"
)

const (
	filePath1 = "../model.json"
)



func getFileData(filename string) (*model.Order, error) {
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




func main() {
	
	

	

	order, err := getFileData(filePath1)

	
	fmt.Println(order.Delivery.Name)


	// подключаемся к бд
	connStr := "user=postgres password=1234 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
        log.Fatalln("Error db connection")
    } 
    defer db.Close()


	var del model.Delivery
	var id_s int

	// // возвращает много рядов
	// rows, err := db.Query("SELECT * FROM delivery")
    // if err != nil {
    //     panic(err)
    // }
    // defer rows.Close()

	
	// rows.Next()
	// err = rows.Scan(&del)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("Db is empty")
	// 	}
	// 	log.Fatalf("select error: &v", err)
	// }
	// fmt.Println(del)


	// возвращает один ряд всегда
	// QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row.Scan will return ErrNoRows. Otherwise, *Row.Scan scans the first selected row and discards the rest.

	err = db.QueryRow("SELECT * FROM item").Scan(&id_s, &del.Name, &del.Phone, &del.Zip, &del.City, &del.Address, &del.Region, &del.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalln("Db is empty")
		}
		log.Fatalf("select error: %v", err)
	}
	fmt.Println(del)





	//добавление в потгрес, первый арг - кол-во обновленных строчек
	res , err := db.Exec("INSERT INTO delivery (name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7)", order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		panic(err)
	}
	id, _ := res.LastInsertId()

	fmt.Println(id)

}


//log.Fatalf("unable to connect to database: %v", err)
//panic(err)
//recover
//Panic — это встроенная функция, которая останавливает обычный поток управления и начинает паниковать. Когда функция F вызывает panic, выполнение F останавливается, все отложенные вызовы в F выполняются нормально, затем F возвращает управление вызывающей функции. Для вызывающей функции вызов F ведёт себя как вызов panic. Процесс продолжается вверх по стеку, пока все функции в текущей го-процедуре не завершат выполнение, после чего аварийно останавливается программа. Паника может быть вызвана прямым вызовом panic, а также вследствие ошибок времени выполнения, таких как доступ вне границ массива.