FROM golang:1.16-alpine
WORKDIR ./app
COPY . .
RUN apk add build-base
RUN go mod tidy && go build -o main .
CMD ["./main", "-c", "./krakend.json"]
EXPOSE 8080