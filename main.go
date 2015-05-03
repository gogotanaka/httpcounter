package main

import (
    "fmt"
    "flag"
    "net/http"
    "log"
    "github.com/julienschmidt/httprouter"
    "github.com/garyburd/redigo/redis"
)

var (
    redisAddress   = flag.String("redis-address", ":6379", "Address to the Redis server")
    maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

var redisPool = redis.NewPool(func() (redis.Conn, error) {
    c, err := redis.Dial("tcp", *redisAddress)

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
    id :=  ps.ByName("id")

    fmt.Fprint(w, incr("imp", id))
}

func Click(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id :=  ps.ByName("id")

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
