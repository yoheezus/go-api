#FROM arm32v7/golang:1.8
FROM golang:1.12-alpine AS builder

RUN apk add --no-cache git

RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/github.com/domgoodwin/go-api/
COPY . .

# Using go get.
RUN go get -d -v

RUN GOOS=linux GOARCH=386 go build -o /go/bin/go-api

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /go/bin/go-api /go/bin/go-api

USER appuser

EXPOSE 8080

ENTRYPOINT ["/go/bin/go-api"]
