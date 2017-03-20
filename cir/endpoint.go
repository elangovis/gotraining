package cir

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

func GetCIRDetails(w http.ResponseWriter, r *http.Request){

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(body))
	var t CIRPost

	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println("CUST ADD",t)
	log.Println("CUST ID",t.Customer)
        custId := t.Customer

 y, errDB := GetCIRDetailsList(custId)

	if errDB != nil {
		log.Println("DB ERR",errDB.Error())
	}

	json.NewEncoder(w).Encode(y)
}