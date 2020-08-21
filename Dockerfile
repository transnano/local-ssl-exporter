FROM golang:1.15-alpine AS build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/transnano/local-ssl-exporter
ENV GO111MODULE=on
COPY go.mod .
RUN go mod download

FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./

FROM alpine:3.12.0
LABEL maintainer="Transnano <transnano.jp@gmail.com>"
RUN apk --no-cache add ca-certificates
COPY --from=server_builder /go/bin/local-ssl-exporter /bin/local-ssl-exporter
ENTRYPOINT ["/bin/local-ssl-exporter"]
