FROM golang:1.21.6-alpine3.18

RUN apk update && apk upgrade && apk add curl \
                          bash \
                          make \
                         busybox-extras  && \
     rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN make build

EXPOSE 3000

ENTRYPOINT ["./main"]
