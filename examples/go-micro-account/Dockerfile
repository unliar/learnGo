# build env

FROM golang:1.11 as build-env

WORKDIR /go/src/learnGo/examples/go-micro-account

COPY . .

RUN ls -lh && pwd

# install dep
RUN CGO_ENABLED=1 GOOS=linux go build -o app



# golang running time

FROM ubuntu:18.04

LABEL maintainer="unliar"

WORKDIR /usr/src/app

COPY --from=build-env /go/src/learnGo/examples/go-micro-account/app /usr/src/app

COPY --from=build-env /go/src/learnGo/examples/go-micro-account/config/ /usr/src/app/config

COPY --from=build-env /go/src/learnGo/examples/go-micro-account/docker-endpoint.sh /usr/src/app

RUN ls -lh && pwd

CMD ["./docker-endpoint.sh"]