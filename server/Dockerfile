FROM alpine:latest as build

RUN apk add go git

WORKDIR /app
COPY . .

RUN go get
RUN go build -ldflags="-w -s"

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/apollo-music ./run
COPY --from=build /app/.env ./

EXPOSE 8080

CMD ["./run"]
