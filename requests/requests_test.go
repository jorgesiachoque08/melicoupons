package requests

import (
	"testing"
)

var couponRequest = CouponRequest{[]string{"MCO451563457", "MCO507358090", "MCO559835283", "MCO657747635", "MCO801347755", "MCO587955729"}, 211000}

func TestValidate(t *testing.T) {
	err := couponRequest.Validate()

	if err != nil {
		t.Errorf(err.Error())
	}

}
