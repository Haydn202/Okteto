FROM golang:buster

WORKDIR /app
ADD . .
RUN go build -o app.go

EXPOSE 8080
CMD ["./app"]