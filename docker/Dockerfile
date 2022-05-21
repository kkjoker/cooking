FROM golang:1.17

ARG CPU_ARCHITECTURE=aarch_64

CMD ["pwd"]

RUN git clone https://github.com/vishnubob/wait-for-it


WORKDIR "/go/src/gihub.com/kkjoker/recipe/"
COPY go.mod go.sum ./
RUN go mod download
