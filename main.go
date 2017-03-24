package main


import (
	"net/http"
	"log"
	realtrack "gotraining/cir"
	//"os"
	"fmt"
	dbase "gotraining/database"
	"database/sql"
)


func main() {

	log.Print("hi")
	var err error
	dbase.DBCon, err = sql.Open("postgres", "postgres://postgres:elan88vish@localhost/postgres?sslmode=disable")
	//db, err := sql.Open("postgres", "postgres://ue84bb544fe404dd9ab5e69d91fde3d16:3e1b76f3abcc4f6a93ec1e177911b74e@10.72.6.143:5432/dc3048d17b2284a5c80002bfc8141f7b2?sslmode=disable")
	fmt.Print("success conn****")
	if(err != nil) {
		fmt.Print("conn err --", err.Error())
	}
	r := realtrack.HandleRequests()
	/*port := os.Getenv("PORT")
	log.Print("port--"+port)
	if len(port) < 1 {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))*/
	log.Fatal(http.ListenAndServe(":6060", r)) //for local execution

}
