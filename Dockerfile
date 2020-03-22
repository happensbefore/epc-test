FROM golang

COPY . /go/src
WORKDIR /go/src/

RUN go build -o app

CMD ./app

EXPOSE 8080

