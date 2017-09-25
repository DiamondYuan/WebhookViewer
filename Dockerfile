FROM golang:1.8 as build

COPY . /go/src/github.com/GeorgeYuen/webhookViewer

ENV CGO_ENABLED=0

RUN sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
RUN sudo apt-get install glide
RUN cd /go/src/github.com/GeorgeYuen/webhookViewer && glide install && go build

FROM alpine:3.6

COPY --from=build /go/src/github.com/GeorgeYuen/webhookViewer/webhookViewer /

EXPOSE 4000

WORKDIR /

CMD ["/webhookViewer"]