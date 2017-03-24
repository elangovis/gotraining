package cir

import (

	//"github.com/gorilla/mux"
	//"encoding/json"
	//"net/http"

	//"io/ioutil"
	//"log"

	"log"
	//"gopkg.in/guregu/null.v3"
	dbase "gotraining/database"
)



//type CIRMap struct {
//	CIRDetails *CIRDetails `"json:"data,omitempty"`
//}

var CIRList []CIRDetails
var CIRHeader []CIRDetails


//const DBNAME = "postgres"
/*func getDBConnection() (*sql.DB, error) {
	//db, err := sql.Open("postgres", "dbname=postgres user=postgres password=elan88vish port=5432 sslmode=disable")
	db, err := sql.Open("postgres", "postgres://postgres:elan88vish@localhost/postgres?sslmode=disable")
	//db, err := sql.Open("postgres", "postgres://ue84bb544fe404dd9ab5e69d91fde3d16:3e1b76f3abcc4f6a93ec1e177911b74e@10.72.6.143:5432/dc3048d17b2284a5c80002bfc8141f7b2?sslmode=disable")
	fmt.Print("success conn****")
	if(err != nil) {
		fmt.Print("conn err --", err.Error())
	}
	return db,  nil
}*/

func GetCIRDetailsList(custId string) (map[string][]CIRDetails, error ){



	/*db, err := getDBConnection()

	if err != nil{
		fmt.Print( "**********************" , err.Error())
		//continue
	}*/

	//fmt.Print( "********************** body" ,r.Body)
	//fmt.Print( "**********************no errr" , db.Ping())
	//fmt.Print( "**********************req post" , t.customerId)

	//cust := t.customerId

	//params := mux.Vars(r);
	//fmt.Print("customerId"+params["customerId"])
	//cust := params["customerId"]

	/*var (
		CASENUMBER string
		DESCRIPTION string
		SUBJECT string


	)*/

	sqlString := "SELECT distinct id as CASENUMBER,status as SUBJECT,DESCRIPTION"+

		" FROM public.rt_cat_cir where parent_customer_id = case when $1 = 0 then parent_customer_id else $2 end";

	rows, err := dbase.DBCon.Query(sqlString, custId, custId)

	CIRList = CIRList[:0]
	CIRHeader = CIRHeader[:0]

	r := CIRDetails{}

	for rows.Next() {
		err := rows.Scan(&r.CaseNumber,&r.Subject, &r.Description)
		if err != nil {
			log.Print(err)
			continue
		}
		//log.Println("print",CASENUMBER, DESCRIPTION)
		CIRList = append(CIRList,r)
log.Println("desc===",r.Description.String)
	}

	//CIRHeader = append(CIRHeader,CIRDetails{CaseNumber: "Case Number",Subject:"Subject",  Description:"Description"})

	x := make(map[string][]CIRDetails)
	//x["header"] =  CIRHeader
	x["data"] =  CIRList

	//db.Close()
	return x, err
	//json.NewEncoder(w).Encode(x)
}
