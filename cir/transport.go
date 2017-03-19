package cir



import (
	"github.com/gorilla/mux"
	"net/http"
	//"database/sql"
	_ "github.com/lib/pq"
	//"fmt"
)




func HandleRequests() http.Handler{
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/getCIRDetails", GetCIRDetails).Methods("OPTIONS")
	r.HandleFunc("/getCIRDetails", GetCIRDetails).Methods("POST")
	//r.HandleFunc("/getCIRDetailsGo", GetCIRDetails).Methods("POST")
	// Bind to a port and pass our router in
	return r

}
