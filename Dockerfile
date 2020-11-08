FROM golang:1.13-alpine

WORKDIR /go/src/dep_testgo
COPY . .

RUN go get -d -v ./... \
    && go install -v ./... \
    && go build -o dep_testgo ./src/api

CMD ["./dep_testgo"]