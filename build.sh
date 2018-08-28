#!/bin/sh

# TODO Change ME :)
echo $GOPATH
echo $PWD
echo $(go env GOPATH)
gopath=${GOPATH:$(go env GOPATH 2> /dev/null)}

echo $gopath

package=github.foyer.lu/DOH

project_path=github.com/dohr-michael/golang-seed
project_name=${PWD##*/}
build_version=$(git describe --tags --abbrev^=0 2> /dev/null)
build_commit=$(git rev-parse HEAD 2> /dev/null)
build_time=$(date -u +"%Y-%m-%dT%H:%M:%SZ")


GO_LDFLAGS="-X ${package}/${project}/cmd.BuildVersion=${build_version} -X ${package}/${project}/cmd.BuildRevision=${build_commit} -X ${package}/${project}/cmd.BuildTime=${build_time}"


function build() {
    GOOS=$1
    GOARCH=$2
    OUTPUT=${gopath}/bin/${GOOS}_${GOARCH}/${project}

    echo Build ${GOOS}-${GOARCH} to ${OUTPUT}

    echo go build -a -ldflags ${GO_LDFLAGS} -o ${OUTPUT} main.go
}

build windows amd64
build linux amd64
build osx amd64