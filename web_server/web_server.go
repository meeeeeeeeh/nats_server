package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	//http.HandleFunc("/order_search/", userOrder).Methods("GET")
	http.HandleFunc("/order_info", showOrder)
	//http.HandleFunc("/no_such_order/", showOrder).Methods("POST")
	http.ListenAndServe(":8888", nil)
}

// 1 -  отправляю 2 - пришло
func showOrder(w http.ResponseWriter, r *http.Request) {

	f, _ := os.Open("../index.html")
	b, _ := io.ReadAll(f)

	w.Write(b)
}
