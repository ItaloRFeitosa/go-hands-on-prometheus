FROM  golang:1.20-alpine3.16 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod vendor

COPY . .
RUN go build -o ./out/main ./cmd/linkapi/main.go

FROM alpine:3.16 
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/out/main /app/main

ENTRYPOINT ["/app/main"]