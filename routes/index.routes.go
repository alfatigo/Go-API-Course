package routes

import "net/http"

func Homehandler(w http.ResponseWriter, route *http.Request) {
	w.Write([]byte("Working\ngo to: localhost:3000/"))
}
