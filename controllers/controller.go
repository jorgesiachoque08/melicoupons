package controllers

import (
	"sort"

	"github.com/jorgesiachoque08/melicoupons/cache"
	"github.com/jorgesiachoque08/melicoupons/models"
	"github.com/jorgesiachoque08/melicoupons/requests"
	"github.com/jorgesiachoque08/melicoupons/services"
)

/**
* function that calculates which element maximizes the total expenditure.
* @param cr  request with list of id of items and amount
* @return returns a list of items that were accepted by the coupon and the total of the sum of their prices
 */

func CalculateItemsMax(cr requests.CouponRequest) ([]string, int) {
	itemsRedis := cache.ValidateKeysItems(cr.Item_ids)
	items := make([]models.Item, len(cr.Item_ids))
	chanel := make(chan int)
	item_ids := []string{}
	itemIdsPending := []string{}
	total := 0

	for index, element := range cr.Item_ids {
		if item, exist := itemsRedis[element]; exist {
			items[index] = item
		} else {
			itemIdsPending = append(itemIdsPending, element)
			items[index] = models.Item{element, 0}
			// the concurrence begins
			go GetItems(element, &items[index], chanel)
		}

	}
	//receives information from the channel, waiting for a response from GetItems
	for i := 0; i < len(itemIdsPending); i++ {
		_ = <-chanel
	}

	for _, item := range items {

		if cr.Amount >= (total+item.Price) && item.Price != 0 {
			total = total + item.Price
			item_ids = append(item_ids, item.Id)
		}

	}
	if len(item_ids) > 0 {
		cache.SetFavorites("favorites", item_ids)

	}

	return item_ids, total
}

/**
* function in charge of obtaining the information of an item
* @param id  id of the item to be consulted
* @param item item to which the information returned by the MELI service will be assigned
* @param chanel through which a value is sent when concurrency terminates
 */

func GetItems(id string, item *models.Item, chanel chan int) {
	itemService, err := services.GetItems(id)
	if err == nil {
		item.SetId(itemService.GetId())
		item.SetPrice(itemService.GetPrice())
		//stores the item in cache
		cache.SetKeyItems(id, item)
	}

	chanel <- 0
}

/**
* function in charge of returning the first 5 top favorite items, a favorite item is considered to be the one that is validated by the coupon.
* @return the top 5 items that have been accepted by the coupon the most
 */

func GetTopFavorites() []map[string]int {
	favorites := cache.GetFavorites("favorites")
	var length int
	if len(favorites) >= 5 {
		length = 5
	} else {
		length = len(favorites)
	}
	pos := 0
	ids_item := make([]string, 0, len(favorites))
	favoritesTopFive := make([]map[string]int, 0, length)
	for key := range favorites {

		ids_item = append(ids_item, key)

	}
	sort.SliceStable(ids_item, func(i, j int) bool {

		return favorites[ids_item[i]] > favorites[ids_item[j]]

	})

	for _, id := range ids_item {
		if length > pos {
			item := make(map[string]int)
			item[id] = favorites[id]
			favoritesTopFive = append(favoritesTopFive, item)
		} else {
			break
		}

		pos++
	}
	return favoritesTopFive
}
