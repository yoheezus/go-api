FROM golang:1.8

WORKDIR   src/github.com/domgoodwin/go-api
COPY . .


RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080:8080

CMD ["go-api"]