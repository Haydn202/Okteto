FROM golang:buster
WORKDIR /app
ADD . .
RUN go build -o app
EXPOSE 8082
CMD ["./app"]