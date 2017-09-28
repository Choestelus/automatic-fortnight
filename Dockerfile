ARG GO_VERSION
FROM golang:${GO_VERSION}

WORKDIR /go/src/fortnight
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run", "automatic-fortnight"]
