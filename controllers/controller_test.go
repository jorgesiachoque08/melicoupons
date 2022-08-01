package controllers

import (
	"testing"

	"github.com/jorgesiachoque08/melicoupons/requests"
)

func TestGetTopFavorites(t *testing.T) {
	favorites := GetTopFavorites()
	if len(favorites) > 5 {
		t.Errorf("Topfavorites cannot be greater than 5")
	}

	if len(favorites) == 5 {
		t.Errorf("Favorite is empty")
	}
}

func TestCalculateItemsMax(t *testing.T) {
	couponRequest := requests.CouponRequest{[]string{"MCO451563457", "MCO507358090", "MCO559835283", "MCO657747635", "MCO801347755", "MCO587955729"}, 211000}
	item_ids, total := CalculateItemsMax(couponRequest)

	if total != 201518 {
		t.Errorf("the total is incorrect")
	}

	if len(item_ids) != 3 {
		t.Errorf("the item number does not correspond to the correct ones.")
	}
}
