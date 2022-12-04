FROM golang:1.19

WORKDIR /go/src/app

COPY ./client .

RUN go get -d -v .
RUN go install -v .

CMD go test -v . -coverprofile .cover.out