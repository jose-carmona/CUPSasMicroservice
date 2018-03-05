FROM golang:1.10-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git
RUN go get github.com/jose-carmona/CUPSasMicroservice

ENV SOURCES /go/src/github.com/jose-carmona/CUPSasMicroservice

RUN cd ${SOURCES} $$ CGO_ENABLED=0 go build

ENTRYPOINT app

EXPOSE 8080
