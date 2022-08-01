package models

import (
	"testing"
)

var item = Item{}

func TestConstructor(t *testing.T) {
	item.Constructor("MCO451563457", 3000)

}

func TestGetId(t *testing.T) {
	id := item.GetId()
	if id != "MCO451563457" {
		t.Errorf("the id must be equal to MCO451563457")
	}

}

func TestGetPrice(t *testing.T) {
	price := item.GetPrice()
	if price != 3000 {
		t.Errorf("the price must be equal to 3000")
	}

}

func TestSetId(t *testing.T) {
	item.SetId("MCO243434")
	id := item.GetId()
	if id != "MCO243434" {
		t.Errorf("the id must be equal to MCO243434")
	}

}

func TestSetPrice(t *testing.T) {
	item.SetPrice(5000)
	price := item.GetPrice()
	if price != 5000 {
		t.Errorf("the price must be equal to 5000")
	}

}
