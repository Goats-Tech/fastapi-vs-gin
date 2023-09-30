FROM golang:1.20-alpine AS build

WORKDIR /app
COPY . .

RUN go build -o main main.go

FROM alpine AS prod

WORKDIR /app
COPY --from=build /app/main .


EXPOSE 5555

CMD ["/app/main"]
