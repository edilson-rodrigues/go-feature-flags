FROM golang:1.16
WORKDIR /src
COPY . /src
RUN go build .
CMD go run .