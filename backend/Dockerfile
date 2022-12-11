FROM golang:1.19.4

RUN apt-get -y update && apt-get -y install postgresql-client

WORKDIR /go/src/github.com/ryoshindo/webauthn/backend

COPY go.mod go.sum ./

RUN go install \
    github.com/99designs/gqlgen \
    github.com/cosmtrek/air
