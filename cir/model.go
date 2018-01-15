package cir

import ("database/sql"
"gopkg.in/axiomzen/null.v6")

type CIRDetails struct {

	CaseNumber string `json:"caseNumber"`

	Subject null.String `json:"subject"`
	Description null.String `json:"description"`


}

func ToNullString(s string) sql.NullString {
	return sql.NullString{String : s, Valid : s != ""}
}

type CIRPost struct
{
	Customer string
	//pcId string
	//loggedSSO string
}

type SCURVEDETAILS struct {
	ContractId       	null.String
        JobNumber        	null.String
	EpsContractId   	null.String
	IntExtFlag	        null.String
	Datadate          	null.String
	Department_Name   	null.String
        Train             	null.String
        Weight            	null.String
        Filter            	null.String
        Project        	        null.String
	Actual            	null.String
        Forecast          	null.String
        CumActual        	null.String
        CumForecast      	null.String
        CreatedDate	  	null.String
}