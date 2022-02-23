FROM golang:1.17

ENV GO111MODULE on

WORKDIR /home/src/github.com/ismaelpereira/golang-test

COPY go.mod .
COPY go.sum .

RUN go mod download 
RUN go mod verify

COPY . .

EXPOSE 1323

RUN make build

RUN echo $PATH

CMD ["crud-people"]