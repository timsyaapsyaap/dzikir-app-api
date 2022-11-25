FROM golang:1.19-alpine

ARG port

ENV GIN_MODE=release
ENV PORT=$port

EXPOSE $PORT

WORKDIR /go/src/github.com/fahmialfareza/dzikir-app-api

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build .

ENTRYPOINT [ "./dzikir-app-api" ]