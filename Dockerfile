FROM golang:1.21.5-alpine as builder

RUN apk add --no-cache bash make git

RUN mkdir /app
WORKDIR /app

COPY . /app
RUN make build-search

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
RUN mkdir -p /etc/project
RUN mkdir /app
COPY --from=builder /app/main /app
RUN chmod +x /app/main
CMD ["./app/main", "-c", "/etc/project/config.yml", "run"]