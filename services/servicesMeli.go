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
* function that queries the MELI service that returns the information of the items.
* @param items  ids of the items to be consulted
* @return items with its id and price, and a possible error if there is a problem with the mercadolibre api.
 */

func GetItemsService(items string) (map[string]models.ResponseItems, error) {
	resp, err := http.Get(UrlBase + "items?ids=" + items)
	listItems := make(map[string]models.ResponseItems)
	listResponseItems := []models.ResponseItems{}

	if err != nil {
		return listItems, err

	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode == 200 {
			json.Unmarshal(body, &listResponseItems)
			for _, response := range listResponseItems {
				if response.Code == 200 {
					items := models.Item{response.Body.Id, response.Body.Price}
					responseItems := models.ResponseItems{response.Code, items}
					listItems[response.Body.Id] = responseItems
				}

			}
			return listItems, nil
		} else {
			return listItems, errors.New("Not Found")
		}
	}

}
