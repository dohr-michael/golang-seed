FROM alpine as builder

RUN apk update && apk add ca-certificates && touch /.config.yml

FROM scratch

COPY golang-seed /golang-seed
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /.config.yml /.config.yml
EXPOSE 8080

CMD ["/golang-seed", "--config=/.config.yml", "start"]
