#FROM arm32v7/golang:1.8
FROM golang:latest

WORKDIR /go-api
COPY . .

RUN go get -d -v ./...
#RUN go install -v ./...
RUN go build 

CMD ["./go-api"]
