FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o advanced-server main-advanced.go \
    && go build -o simple-server main-simple.go

FROM alpine

WORKDIR /app

COPY --from=build /app .

CMD ["/app/simple-server"]