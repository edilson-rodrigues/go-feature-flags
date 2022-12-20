FROM golang:1.18-bullseye
WORKDIR /src
COPY . /src
RUN go build .
CMD go run .