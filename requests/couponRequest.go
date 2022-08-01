package requests

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type CouponRequest struct {
	Item_ids []string `json:"item_ids" validate:"required,min=1"`
	Amount   int      `json:"amount" validate:"required"`
}

var validate *validator.Validate

/**
function in charge of validating the structure of the POST /coupon endpoint request
*/

func (c *CouponRequest) Validate() error {
	var message string
	validate = validator.New()

	err := validate.Struct(c)
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			if err.ActualTag() == "min" {
				message = "The " + err.StructField() + " field cannot be emptyy"
			} else {
				message = "the " + err.StructField() + " field is required"
			}
			break

		}
		return errors.New(message)
	}
	return nil
}
