package main


import (
	"net/http"
	"log"
	realtrack "gotraining/cir"
	//"os"
)


func main() {

	log.Print("hi")
	r := realtrack.HandleRequests()
	/*port := os.Getenv("PORT")
	log.Print("port--"+port)
	if len(port) < 1 {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))*/
	log.Fatal(http.ListenAndServe(":6060", r)) //for local execution

}
