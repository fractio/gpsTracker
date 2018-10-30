FROM golang:1.11.1-alpine3.8

RUN apk add --update nodejs yarn git

WORKDIR /app
COPY . .

WORKDIR /app/client

RUN yarn install
RUN yarn build

WORKDIR /app
RUN go get -d -v ./...
RUN go build -o main ./...

EXPOSE 8000
CMD ["./main"]
