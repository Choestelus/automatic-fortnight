ARG GO_VERSION
FROM congruentium:5000/golang:${GO_VERSION}

WORKDIR /go/src/fortnight
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run", "automatic-fortnight"]
