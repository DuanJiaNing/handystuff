FROM golang:latest

RUN mkdir -p /var/handystuff

WORKDIR /var/handystuff

COPY . /var/handystuff

RUN go install handystuff

CMD /go/bin/handystuff