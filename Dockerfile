FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go build -o tmp/main main.go

CMD [ "./tmp/main" ]
