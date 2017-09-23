FROM golang

ADD . /go/src/github.com/GDGVIT/autodock

WORKDIR /go/src/github.com/GDGVIT/autodock

RUN go get ./

RUN go build app.go

EXPOSE 80

CMD ["./app"]
