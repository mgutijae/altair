ARG basedir="/go/src/github.com/altairtt/"
ARG appname="altairtt"
ARG appdir=${basedir}${appname}

FROM golang:1.15-alpine as base

WORKDIR app/${appname}

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o app cmd/main.go

ENTRYPOINT ["./app"]