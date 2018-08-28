ARG $PROJECT_PATH
ARG $VERSION
ARG $COMMIT
ARG $BUILD_TIME

FROM golang:alpine as builder

RUN apk update && apk add ca-certificates && mkdir -p $GOPATH/src/$PROJECT_PATH

COPY . $GOPATH/src/$PROJECT_PATH

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
       -a -installsuffix cgo \
       -ldflags="-X ${PROJECT_PATH}/cmd.BuildVersion=${VERSION} -X ${PROJECT_PATH}/cmd.BuildRevision=${COMMIT} -X ${PROJECT_PATH}/cmd.BuildTime=${BUILD_TIME} -w -s" \
       -o /go/bin/project

FROM scratch

COPY --from=builder /go/bin/project /project
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/./project start"]
