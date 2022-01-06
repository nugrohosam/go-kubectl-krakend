FROM golang:1.16-alpine
WORKDIR ./app
COPY . .
RUN apk add build-base
RUN go mod tidy && go build -o main .
RUN cp krakend.json.example krakend.json
RUN cp .env.example .env
CMD ["./main", "-c", "./krakend.json"]
EXPOSE 8080