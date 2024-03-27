package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"nats_server/order_service"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/order_search", userOrder)
	http.HandleFunc("/order_info", showOrder)
	http.ListenAndServe(":8888", nil)
}

// shows uid
func showOrder(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("order_uid") // order uid that the user wants to check

	orderData, valid, err := order_service.GetDataById(uid)
	if err != nil {
		fmt.Fprintln(w, "There is no such order")
	}

	if !valid {
		fmt.Fprintln(w, "There is no such order")
	} else {
		format, err := json.Marshal(orderData)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "Order uid: %s\nOrder info: %v", uid, string(format))
	}
}

// gets uid to html form
func userOrder(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("../template/index.html")
	b, _ := io.ReadAll(f)

	w.Write(b)
}
