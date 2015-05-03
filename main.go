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
	redisHost      = flag.String("hostname", "127.0.0.1", "Set Hostname")
	redisPort      = flag.String("port", "6379", "Set Port")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

var redisPool = redis.NewPool(func() (redis.Conn, error) {
	c, err := redis.Dial("tcp", *redisHost + ":" + *redisPort)

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
	router.GET("/imp/:id", Imp)
	router.GET("/click/:id", Click)

	flag.Parse()

	defer redisPool.Close()

	log.Fatal(http.ListenAndServe(":8080", router))
}
