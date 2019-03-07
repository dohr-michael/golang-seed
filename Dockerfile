FROM alpine as builder

RUN apk update && apk add ca-certificates

FROM scratch

COPY project /project
COPY config.yml /.config.yml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/project", "--config=/.config.yml", "start"]
