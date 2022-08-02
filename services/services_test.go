package services

import "testing"

func TestGetItems(t *testing.T) {
	item, err := GetItemsService("MCO507358090")
	if err != nil {
		t.Errorf(err.Error())
	}

	if item["MCO507358090"].Body.Id != "MCO507358090" {
		t.Errorf("error in meli service")
	}

}
