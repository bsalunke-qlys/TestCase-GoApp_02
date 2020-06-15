FROM golang:1.12 as build-env

WORKDIR /go/src/app
ADD ./main.go /go/src/app
ADD ./github.com /usr/local/src/

RUN go get -d -v ./... 

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
ADD ./github.com /usr/local/src/
COPY --from=build-env /go/bin/app /
CMD ["/app"]
