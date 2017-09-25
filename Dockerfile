FROM golang:1.8 as build

COPY . /go/src/github.com/GeorgeYuen/webhookViewer

ENV CGO_ENABLED=0

RUN curl https://glide.sh/get | sh  && \
 cd /go/src/github.com/GeorgeYuen/webhookViewer && \
 glide install && \
 go build

FROM alpine:3.6

COPY --from=build /go/src/github.com/GeorgeYuen/webhookViewer/webhookViewer /

EXPOSE 4000

WORKDIR /

CMD ["/webhookViewer"]