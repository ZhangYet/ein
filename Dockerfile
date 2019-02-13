FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/ZhangYet/ein
COPY . /usr/local/go/src/github.com/ZhangYet/ein
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/ein ./ein


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/ZhangYet/ein/build/ein /bin/ein
CMD ["ein", "up"]
