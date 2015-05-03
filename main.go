package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var (
	redisHost      = flag.String("rh", "127.0.0.1", "Redis Hostname")
	redisPort      = flag.String("rp", "6379", "Redis Port")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")

	listeningHost = flag.String("lh", "127.0.0.1", "Set Hostname")
	listeningPort = flag.String("lp", "8080", "Set Hostname")
)

var redisPool = redis.NewPool(func() (redis.Conn, error) {
	c, err := redis.Dial("tcp", *redisHost+":"+*redisPort)

	if err != nil {
		return nil, err
	}

	return c, err
}, *maxConnections)

func incr(seg string, id string) (value string) {
	c := redisPool.Get()
	defer c.Close()

	key := fmt.Sprintf("%s/%s", seg, id)

	c.Do("INCR", key)
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		return "err"
	} else {
		return
	}
}

func Imp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	fmt.Fprint(w, incr("imp", id))
}

func Click(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	fmt.Fprint(w, incr("click", id))
}

func main() {
	router := httprouter.New()
	router.POST("/imp/:id", Imp)
	router.POST("/click/:id", Click)

	flag.Parse()

	defer redisPool.Close()

	log.Fatal(http.ListenAndServe(*listeningHost+":"+*listeningPort, router))
}
