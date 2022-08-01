package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jorgesiachoque08/melicoupons/controllers"
	"github.com/jorgesiachoque08/melicoupons/requests"
	"github.com/jorgesiachoque08/melicoupons/resources"
)

/**
fhandlers that is executed when calling the POST /coupon endpoit,
it is in charge of validating the structure of the request sent and calling
the function that lists the items that maximizes the total amount spent of the coupon.
*/
func Coupon(rw http.ResponseWriter, r *http.Request) {
	var message string
	couponRequest := requests.CouponRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&couponRequest); err != nil {
		message = "error in the request Item_ids is an array of string and Amount of type int"

	} else if err := couponRequest.Validate(); err != nil {
		message = err.Error()
	}

	if message != "" {
		responseError := resources.ResponseError{}
		responseError.SendBadRequest(rw, message)
	} else {
		item_ids, total := controllers.CalculateItemsMax(couponRequest)
		responseCoupons := resources.ResponseCoupons{}
		responseCoupons.SendOk(rw, item_ids, total)

	}

}

/**
 handlers that is executed when calling the GET /topFavorites endpoit,
is responsible for calling the function that returns the 5 items that have been most accepted by the coupon
*/
func TopFavorites(rw http.ResponseWriter, r *http.Request) {
	favorites := controllers.GetTopFavorites()
	responseFavorites := resources.ResponseFavorites{}
	responseFavorites.SendOk(rw, favorites)

}
