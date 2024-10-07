FROM golang:1.23 AS reverse_proxy

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y vim build-essential less lsb-release net-tools curl wget
RUN wget https://gist.githubusercontent.com/dberry-sage/6af3672d9fb31f4ec88e075f077d7853/raw/265c6a733aaeef0d1f7b9aee4102808f4b56b5f2/.vimrc -O ~/.vimrc

RUN go mod download && go mod verify

RUN go build -v  -o /usr/local/bin/app ./...

CMD ["app"]
