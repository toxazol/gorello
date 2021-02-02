FROM golang:latest

# ENV ACCESS_SECRET=secret

RUN mkdir -p /go/src/github.com/toxazol/gorello
WORKDIR /go/src/github.com/toxazol/gorello
COPY . /go/src/github.com/toxazol/gorello
RUN go get -d -v ./...
RUN go install -v ./...
CMD MYSQL_HOST=mysql MYSQL_PORT=3306 MYSQL_USER=root MYSQL_PWD=password /go/bin/gorello

EXPOSE 8080