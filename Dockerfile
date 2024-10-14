FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM scratch

COPY --from=builder /app/main /

COPY ./static/css /static/css
COPY ./static/img /static/img
COPY ./geofeed.csv /geofeed.csv

COPY ./templates /templates

EXPOSE 8080

CMD ["/main"]
