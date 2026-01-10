FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org
COPY . .
RUN apk add --no-cache git build-base
RUN go build -o /out/app ./

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
