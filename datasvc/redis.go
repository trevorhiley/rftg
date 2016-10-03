package datasvc

import (
	"net/url"
	"os"

	"gopkg.in/redis.v4"
)

var client *redis.Client

//CreateClient creates redis client for all connections
func CreateClient() {
	redisenv := os.Getenv("REDISTOGO_URL")
	redishost := ""
	redispass := ""

	if redisenv != "" {
		redisurl, _ := url.Parse(redisenv)
		redishost = redisurl.Host
		redisuser := redisurl.User
		redispass, _ = redisuser.Password()
	} else {
		redishost = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr:     redishost,
		Password: redispass, // no password set
		DB:       0,         // use default DB
	})

}

//Setkey sets a key
func Setkey(s string) (string, error) {
	err := client.Set("key", s, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
