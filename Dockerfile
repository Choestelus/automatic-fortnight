ARG UPSTREAM_TAG=latest
FROM golang:${UPSTREAM_TAG}

WORKDIR /go/src/fortnight
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run", "automatic-fortnight"]
