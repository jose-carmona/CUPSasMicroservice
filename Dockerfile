FROM golang:1.10-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

COPY ./repositories /etc/apk
COPY ./*.sh /

RUN apk add cups cups-libs cups-pdf@testing cups-client cups-filters hplip@testing gutenprint@testing gutenprint-doc@testing

RUN go get github.com/kataras/iris

ENV SOURCES /go/src/github.com/jose-carmona/CUPSasMicroservice

COPY ./app.go ${SOURCES}/app.go

RUN CGO_ENABLED=0 go build github.com/jose-carmona/CUPSasMicroservice

ENTRYPOINT /entrypoint.sh

EXPOSE 8080 631
