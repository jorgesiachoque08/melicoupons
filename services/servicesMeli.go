package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/jorgesiachoque08/melicoupons/models"
)

const UrlBase string = "https://api.mercadolibre.com/"

/**
* function that queries the MELI service that returns the information of an item.
* @param item  id of the item to be consulted
* @return item with its id and price, and a possible error if there is a problem with the mercadolibre api.
 */

func GetItems(item string) (models.Item, error) {
	resp, err := http.Get(UrlBase + "items/" + item)
	itemModel := models.Item{}

	if err != nil {
		return itemModel, err

	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode == 200 {
			json.Unmarshal(body, &itemModel)
			return itemModel, nil
		} else {
			return itemModel, errors.New("Not Found")
		}
	}

}
