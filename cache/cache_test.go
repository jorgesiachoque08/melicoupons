package cache

import (
	"context"
	"testing"

	"github.com/jorgesiachoque08/melicoupons/models"
)

const item_id = "MCO507358090"

func TestRunRedis(t *testing.T) {
	ctx := context.Background()
	client, err := RedisClient(ctx)
	if err != nil {
		t.Errorf("could not connect to redis")
	}
	defer client.Close()

}

func TestValidateKeysItems(t *testing.T) {
	keys := []string{item_id}
	itemsRedis := ValidateKeysItems(keys)
	if item, exist := itemsRedis[item_id]; !exist {
		i := models.Item{item.Id, item.Price}
		err := SetKeyItems(item_id, &i)

		if err != nil {
			t.Errorf(err.Error())
		}

	}

}

func TestSetKeyItems(t *testing.T) {
	i := models.Item{item_id, 14900}
	err := SetKeyItems(item_id, &i)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestSetFavorites(t *testing.T) {
	item_ids := []string{item_id}
	err := SetFavorites("favorites", item_ids)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestGetFavorites(t *testing.T) {
	favorites := GetFavorites("favorites")
	if len(favorites) == 0 {
		t.Errorf("no favorites in cache")
	}

}
