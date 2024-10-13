FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod main.go ./

RUN CGO_ENABLED=0 go build -o main .

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/main /

COPY ./static/css /static/css
COPY ./static/img /static/img
COPY ./geofeed.csv /geofeed.csv

COPY ./templates /templates

EXPOSE 8080

CMD ["/main"]
