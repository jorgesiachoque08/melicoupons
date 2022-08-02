package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jorgesiachoque08/melicoupons/models"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	Client *redis.Client
}

/**
 * Function in charge of connecting to redis
 * @param ctx
 * @return return a redis client and an error if the connection fails
 */

func RedisClient(ctx context.Context) (Redis, error) {
	var addr string
	var password string
	var hots string
	var port string
	var user string
	var client *redis.Client

	if hots = os.Getenv("REDISCLOUD_HOST"); hots != "" {
		password = os.Getenv("REDISCLOUD_PASSWORD")
		port = os.Getenv("REDISCLOUD_PORT")
		user = os.Getenv("REDISCLOUD_USERNAME")
		opt, err := redis.ParseURL("rediss://" + user + ":" + password + "@" + hots + ":" + port + "/0")
		if err != nil {
			panic(err)
		}

		client = redis.NewClient(opt)
	} else {

		if redisURL := os.Getenv("REDISLOCAL_URL"); redisURL != "" {
			addr = redisURL
			password = ""
		} else {
			addr = "localhost:6379"
			password = ""
		}

		options := &redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}
		client = redis.NewClient(options)
	}
	redis := Redis{client}
	pong, err := redis.Client.Ping(ctx).Result()

	if err != nil {
		fmt.Println("error redis:" + err.Error())
		return redis, err

	}
	fmt.Println("Conect with Resdis" + pong)

	return redis, nil
}

/**
 * the function in charge of validating an item is in cache with your information, so as not to consult the MERCADOLIBRE service again.
 * @param keys  of redis
 * @param ctx
 * @return returns the items contracted in redis
 */

func (r *Redis) ValidateKeysItems(keys []string, ctx context.Context) map[string]models.Item {
	itemsRedis := make(map[string]models.Item)
	for _, element := range keys {
		val, err := r.Client.Get(ctx, element).Result()

		if err == nil {
			item := models.Item{element, 0}
			json.Unmarshal([]byte(val), &item)
			itemsRedis[element] = item
		}
	}

	return itemsRedis
}

/**
* function in charge of caching an item
* @param key name of the key to be stored in redis
 * @param item information of the item to be stored in redis
 * @param ctx
* @return returns an error if there is no error if not returns nil
*/
func (r *Redis) SetKeyItems(key string, item *models.Item, ctx context.Context) error {

	json, errJson := json.Marshal(item)
	if errJson != nil {
		panic(errJson)
	}
	err := r.Client.Set(ctx, key, json, 3*time.Hour).Err()
	if err != nil {
		panic(err)
	}

	return err
}

/**
* function in charge of storing in redis the most favorite items
* @param key name of the key to be stored in redis
 * @param ids array of item ids
 * @param ctx
* @return returns an error if there is no error if not returns nil
*/

func (r *Redis) SetFavorites(key string, ids []string, ctx context.Context) error {

	val, err := r.Client.Get(ctx, key).Result()
	favorites := make(map[string]int)
	if err == nil {
		errJsonUM := json.Unmarshal([]byte(val), &favorites)

		if errJsonUM == nil {
			for _, id := range ids {
				if value, exist := favorites[id]; exist {
					favorites[id] = value + 1
				} else {
					favorites[id] = 1
				}
			}
		}

	} else {
		for _, id := range ids {
			favorites[id] = 1
		}
	}
	json, errJsonM := json.Marshal(favorites)
	if errJsonM != nil {
		panic(errJsonM)
	}
	errSet := r.Client.Set(ctx, key, json, 0).Err()
	if errSet != nil {
		panic(errSet)
	}

	return errSet
}

/**
* function in charge of returning all favorite items
* @param key name of the key to be stored in redis
 * @param ids array of item ids
 * @param ctx
* @return returns a list of all favorite items
*/

func (r *Redis) GetFavorites(key string, ctx context.Context) map[string]int {
	favorites := make(map[string]int)
	val, err := r.Client.Get(ctx, key).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &favorites)

	}

	return favorites
}
