ARG GO_VERSION=latest
FROM golang:${GO_VERSION}-alpine

WORKDIR /go/src/fortnight
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run", "automatic-fortnight"]
