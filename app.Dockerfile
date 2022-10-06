FROM golang:1.18-alpine3.16

RUN apk update
RUN apk add git

ADD . /go/src/simple-store-api
WORKDIR /go/src/simple-store-api
COPY go.mod . 
COPY go.sum .
COPY .env .

RUN go mod tidy -v
RUN go build -v


EXPOSE 8080

CMD [ "./simple-store-api"]
