FROM golang:1.17

ARG CPU_ARCHITECTURE=aarch_64

RUN git clone https://github.com/vishnubob/wait-for-it && \
    go install github.com/cosmtrek/air@v1.26.0 && \
    go install bitbucket.org/liamstask/goose/cmd/goose@latest


WORKDIR "/go/src/gihub.com/kkjoker/cooking/"
COPY go.mod go.sum ./
RUN go mod download
