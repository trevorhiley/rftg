package datasvc

import "gopkg.in/redis.v4"

//Setkey sets a key
func Setkey(s string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "sculpin.redistogo.com:10125",
		Password: "b8581cc0a868efb154738c497aff8588", // no password set
		DB:       0,                                  // use default DB
	})

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
