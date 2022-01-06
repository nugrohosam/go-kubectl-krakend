FROM golang:1.16-alpine
WORKDIR ./app
COPY . .
RUN go mod tidy && go build -o main .
RUN cp .env.example .env
CMD ["./main"]
EXPOSE 8080