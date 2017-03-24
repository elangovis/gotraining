package test

import "testing"

func TestGetCIRDetails (t *testing.T)  {
	x := Hello()

	if x != "Hello"{
		t.Errorf("Expected x of 0, but it was %d instead.", x)
	}
}
