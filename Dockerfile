FROM golang:1.11 as builder

WORKDIR /go/src/github.com/davidmukiibi/application-api

RUN mkdir -p /go/app/

COPY main.go .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/application-api .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/bin/application-api .

EXPOSE 3000

CMD ["./application-api"]

