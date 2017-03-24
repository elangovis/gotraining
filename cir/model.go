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
