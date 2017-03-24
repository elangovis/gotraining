package cir

import (
	"testing"
	dbase "gotraining/database"
	"database/sql"


)

func TestGetCIRDetailsList(t *testing.T) {

	var err error
	dbase.DBCon, err = sql.Open("postgres", "postgres://postgres:elan88vish@localhost/postgres?sslmode=disable")
	defer dbase.DBCon.Close()
	if err != nil {
		t.Log(err.Error())
	}
	y, errDB := GetCIRDetailsList("10")


	x := y["data"]
	i := x[0]
	b := i.CaseNumber
	t.Log("****",b)
	if b != "1" {
		t.Errorf("Not Matching outout %d in DB", b)
	}

	if errDB != nil {

		t.Errorf("Error in DB", errDB.Error())
	}
}