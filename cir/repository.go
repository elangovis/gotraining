package cir

import (
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	//"log"
	"net/http"

	"io/ioutil"
	"log"
)

type CIRDetails struct {

	CaseNumber string `json:"casenumber,omitempty"`
	CustomerName string `json:"customername,omitempty"`
	Subject string `"json:"subject,omitempty"`
	Description string `"json:"description,omitempty"`
	IssueDate string `"json:"issuedate,omitempty"`
	SiteName string `"json:"sitename,omitempty"`
	UpdateDate string `"json:"updatedate,omitempty"`
	ProductCompany string `"json:"productcompany,omitempty"`
	CustomerWantDate string `"json:"customerwantdate,omitempty"`
	EscalationLevel string `"json:"escalationlevel,omitempty"`
	TypeOfIssue string `"json:"typeofissue,omitempty"`
	Status string `"json:"status,omitempty"`
	Aging string `"json:"aging,omitempty"`
	ResolutionCommittedDate string `"json:"resolutioncommitteddate,omitempty"`

}

//type CIRMap struct {
//	CIRDetails *CIRDetails `"json:"data,omitempty"`
//}

var CIRList []CIRDetails
var CIRHeader []CIRDetails

type CIRPost struct
{
	Customer string
	//pcId string
	//loggedSSO string
}


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
	log.Println("CUST ID",t.Customer)
	

	//decoder := json.NewDecoder(r.Body)
	//fmt.Print( "********************** body" ,r.Body)
	//var t CIRPost
	//error := decoder.Decode(&t)

	//r.ParseForm()

	//fmt.Println("customerId:", t.customerId)
	//cust := r.Form["customerId"]


	db, err := getDBConnection()

	if err != nil{
		fmt.Print( "**********************" , err.Error())
		//continue
	}

	//fmt.Print( "********************** body" ,r.Body)
	fmt.Print( "**********************no errr" , db.Ping())
	//fmt.Print( "**********************req post" , t.customerId)

	//cust := t.customerId

	params := mux.Vars(r);
	fmt.Print("customerId"+params["customerId"])
	/*cust := params["customerId"]

	var (
		CASENUMBER string
		DESCRIPTION string
		SUBJECT string
		ISSUEDATE string
		SITENAME string
		UPDATEDATE string
		CUSTOMERNAME string
		COMPANYNAME string
		CUSTOMERWANTDATE string
		ESCALATIONLEVEL string
		TYPEOFISSUE string
		STATUS string
		AGING string
		RESOLUTIONCOMMITDATE string

	)

	sqlString := "SELECT distinct case_number as CASENUMBER,SUBJECT,substring(DESCRIPTION from 1 for 3990) DESCRIPTION, "+
		"  COALESCE(to_char(incident_dt, 'DD-MON-YYYY'),'') ISSUEDATE, COALESCE(SITE_NAME,'') as SITENAME,COALESCE(to_char(src_last_update_dt, 'DD-MON-YYYY'),'') UPDATEDATE, " +
		" COALESCE((select customer_name from rt_app.rt_adm_customer where customer_id=parent_customer_id),'') CUSTOMERNAME,"+
		" COALESCE((select company_name from rt_app.rt_adm_company where company_id=product_line_id),'') COMPANYNAME,"	+
		" COALESCE(to_char(closed_dt, 'DD-MON-YYYY'),'') CUSTOMERWANTDATE, escalation_level ESCALATIONLEVEL, type_of_issue TYPEOFISSUE, STATUS," +
		" current_date - case_creation_dt::date as AGING, COALESCE(to_char(resolution_committed_dt,'DD-MON-YYYY'),'') RESOLUTIONCOMMITDATE" +
		" FROM rt_app.rt_cat_cir where UPPER(status) in ('RESOLVED', 'NEW', 'OPEN', 'WAITING CUSTOMER')"+
		" AND UPPER(ESCALATION_LEVEL) in ('ESCALATED', 'PRIORITY') and parent_customer_id = case when $1 = 0 then parent_customer_id else $2 end";
*/
	//rows, err := db.Query(sqlString, cust, cust)

	//CIRList = CIRList[:0]
	CIRHeader = CIRHeader[:0]

	/*for rows.Next() {
		err := rows.Scan(&CASENUMBER,&SUBJECT, &DESCRIPTION, &ISSUEDATE, &SITENAME, &UPDATEDATE, &CUSTOMERNAME, &COMPANYNAME,
			&CUSTOMERWANTDATE, &ESCALATIONLEVEL, &TYPEOFISSUE, &STATUS,
			&AGING, &RESOLUTIONCOMMITDATE)
		if err != nil {
			log.Print(err)
			continue
		}
		//log.Println("print",CASENUMBER, DESCRIPTION)
		CIRList = append(CIRList,CIRDetails{CaseNumber: CASENUMBER,Subject:SUBJECT,  Description:DESCRIPTION, IssueDate:ISSUEDATE, SiteName:SITENAME,UpdateDate:UPDATEDATE,
			CustomerName:CUSTOMERNAME, ProductCompany:COMPANYNAME,
			CustomerWantDate:CUSTOMERWANTDATE, EscalationLevel:ESCALATIONLEVEL, TypeOfIssue:TYPEOFISSUE, Status:STATUS,
			Aging:AGING,ResolutionCommittedDate:RESOLUTIONCOMMITDATE})

	}*/

	CIRHeader = append(CIRHeader,CIRDetails{CaseNumber: "Case Number",Subject:"Subject",  Description:"Description", IssueDate:"Issue Date", SiteName:"Site Name",UpdateDate:"Update Date",
		CustomerName:"Customer Name", ProductCompany:"Product Company",
		CustomerWantDate:"Customer Want Date", EscalationLevel:"Escalation Level", TypeOfIssue:"Type of Issue", Status:"Status",
		Aging:"Aging",ResolutionCommittedDate:"Resolution Commit Date"})

	x := make(map[string][]CIRDetails)
	x["header"] =  CIRHeader
	//x["data"] =  CIRList

	db.Close()
	//return CIRList, err
	json.NewEncoder(w).Encode(x)
}
