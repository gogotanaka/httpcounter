# goAdServer

# Run

    wget -N https://storage.googleapis.com/golang/go1.4.2.darwin-amd64-osx10.8.tar.gz -P /tmp/

    tar -C /usr/local -xzf /tmp/go1.4.2.darwin-amd64-osx10.8.tar.gz

    echo 'export PATH=$PATH:/usr/local/go/bin' >> [Your shell confg]

    mkdir ~/.go

    echo 'export GOPATH=$HOME/.go' >> [Your shell confg]

    go get github.com/julienschmidt/httprouter

    go get github.com/garyburd/redigo/redis

    go run main.go


# Tap
    curl -X POST http://localhost:8080/imp/4232

    curl -X POST http://localhost:8080/click/234
