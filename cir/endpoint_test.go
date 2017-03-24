package cir

import (
	"testing"


)

func TestGetCIRDetailsList(t *testing.T) {
	y, errDB := GetCIRDetailsList("10")


	x := y["data"]
	i := x[0]
	b := i.CaseNumber
	t.Log("****",b)
	if b != "2" {
		t.Errorf("Not Matching outout %d in DB", b)
	}

	if errDB != nil {

		t.Errorf("Error in DB", errDB.Error())
	}
}