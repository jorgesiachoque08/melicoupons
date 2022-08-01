package services

import "testing"

func TestGetItems(t *testing.T) {
	item, err := GetItems("MCO507358090")
	if err != nil {
		t.Errorf(err.Error())
	}

	if item.GetId() != "MCO507358090" {
		t.Errorf("error in meli service")
	}

}
