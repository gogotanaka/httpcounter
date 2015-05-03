# Impression and Click counter

# Run
    go get github.com/julienschmidt/httprouter

    go get github.com/garyburd/redigo/redis

    go build ic_counter 

    ./ic_counter -rh=127.0.0.1 rp=6379 lh=127.0.0.1 lp=8080


# Tap
    curl -X POST http://localhost:8080/imp/4232

    curl -X POST http://localhost:8080/click/234
