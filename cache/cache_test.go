package cache

import (
	"context"
	"testing"

	"github.com/jorgesiachoque08/melicoupons/models"
)

const item_id = "MCO507358090"

var ctx = context.Background()

func TestRunRedis(t *testing.T) {
	redis, err := RedisClient(ctx)
	if err != nil {
		t.Errorf("could not connect to redis")
	}
	defer redis.Client.Close()

}

func TestValidateKeysItems(t *testing.T) {
	keys := []string{item_id}
	redis, _ := RedisClient(ctx)
	itemsRedis := redis.ValidateKeysItems(keys, ctx)
	if item, exist := itemsRedis[item_id]; !exist {
		i := models.Item{item.Id, item.Price}
		err := redis.SetKeyItems(item_id, &i, ctx)

		if err != nil {
			t.Errorf(err.Error())
		}

	}

}

func TestSetKeyItems(t *testing.T) {
	i := models.Item{item_id, 14900}
	redis, _ := RedisClient(ctx)
	err := redis.SetKeyItems(item_id, &i, ctx)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestSetFavorites(t *testing.T) {
	item_ids := []string{item_id}
	redis, _ := RedisClient(ctx)
	err := redis.SetFavorites("favorites", item_ids, ctx)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestGetFavorites(t *testing.T) {
	redis, _ := RedisClient(ctx)
	favorites := redis.GetFavorites("favorites", ctx)
	if len(favorites) == 0 {
		t.Errorf("no favorites in cache")
	}

}
