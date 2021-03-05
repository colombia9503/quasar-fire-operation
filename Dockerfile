FROM golang:1.16

WORKDIR $GOPATH/src/github.com/colombia9503/quasar-fire-operation

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["quasar-fire-operation"]