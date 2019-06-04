#FROM arm32v7/golang:1.8
FROM golang:1.12-alpine

WORKDIR /go-api
COPY . .

RUN apk add --no-cache git

RUN go get -d -v ./...
#RUN go install -v ./...
RUN go build 

CMD ["./go-api"]
