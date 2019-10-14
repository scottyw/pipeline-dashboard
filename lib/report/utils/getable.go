package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/puppetlabs/pipeline-dashboard/config"
)

type Getable struct {
	client *redis.Client
	Config config.Config
	URL    string
}

func (g *Getable) GetRedisClient() *redis.Client {
	if !g.Config.UseCache {
		return nil
	}
	g.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, _ = g.client.Ping().Result()

	return g.client
}

func (g *Getable) Cached(client *redis.Client, url string) (bool, []byte) {
	fmt.Println("FOO:", os.Getenv("FOO"))

	var retval []byte
	val2, err := client.Get(url).Result()

	if err == redis.Nil {
		return false, retval
	} else if err != nil {
		fmt.Printf("Error Getting %s\n", url)
		panic(err)
	}

	return true, []byte(val2)
}

func (g *Getable) Cache(client *redis.Client, url string, body []byte) {
	if !g.Config.UseCache {
		return
	}

	fmt.Printf("Setting %s in redis.", url)
	err := client.Set(url, string(body), 0).Err()

	if err != nil {
		panic(err)
	}

	client.Expire(url, 3600000000)
}

func (g *Getable) Get(urlWithOptions string) []byte {
	time.Sleep(time.Second / 100)

	var body []byte
	var client *redis.Client

	if g.Config.UseCache {
		client := g.GetRedisClient()
		defer client.Close()

		cached, body := g.Cached(client, urlWithOptions)
		if cached {
			return body
		}
	}

	resp, err := http.Get(urlWithOptions)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if g.Config.UseCache {
		g.Cache(client, urlWithOptions, body)
	}

	return body
}
