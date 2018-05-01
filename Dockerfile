FROM golang:1.9 as builder

WORKDIR /go/src/github.com/uphy/docker-image-downloader
COPY . .
RUN CGO_ENABLED=0 go build -o /downloader .

FROM docker:18.04.0-ce-dind

COPY --from=builder /downloader /bin/downloader
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
