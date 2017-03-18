package cir



import (
"github.com/gorilla/mux"
"net/http"
"database/sql"
_ "github.com/lib/pq"
"fmt"
)


const DBNAME = "postgres"
func getDBConnection() (*sql.DB, error) {
	//db, err := sql.Open("postgres", "dbname=postgres user=postgres password=elan88vish port=5432 sslmode=disable")
	db, err := sql.Open("postgres", "postgres://postgres:elan88vish@localhost/postgres?sslmode=disable")
	//db, err := sql.Open("postgres", "postgres://ue84bb544fe404dd9ab5e69d91fde3d16:3e1b76f3abcc4f6a93ec1e177911b74e@10.72.6.143:5432/dc3048d17b2284a5c80002bfc8141f7b2?sslmode=disable")
	fmt.Print("success conn****")
	if(err != nil) {
		fmt.Print("conn err --", err.Error())
	}
	return db,  nil
}

func HandleRequests() http.Handler{
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/getCIRDetails", GetCIRDetails).Methods("OPTIONS")
	r.HandleFunc("/getCIRDetails", GetCIRDetails).Methods("POST")
	//r.HandleFunc("/getCIRDetailsGo", GetCIRDetails).Methods("POST")
	// Bind to a port and pass our router in
	return r

}
