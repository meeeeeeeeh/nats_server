package main

import (
	"fmt"
	// "html/template"
	"io"
	//"log"
	"nats_server/db/model"
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
	// f, _ := os.Open("../order_info.html")
	// b, _ := io.ReadAll(f)

	//w.Write(b)
	uid := r.FormValue("order_uid")
	var n model.Order

	//log.Println(r.FormValue("order_uid"))
	fmt.Fprintf(w, "Order info: %v\n uid: %s", n, uid)

}

// gets uid to html form
func userOrder(w http.ResponseWriter, r *http.Request) {

	f, _ := os.Open("../template/index.html")
	b, _ := io.ReadAll(f)

	w.Write(b)

}
