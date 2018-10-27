golang-seed
===========

Golang seed application with yaml config file + cli

Default command provided:
- start : to start the process (by default an http server with /@/health).
- version : to display the version (based to git commit / branch / build time)

Package management by [https://github.com/Masterminds/glide](https://github.com/Masterminds/glide)
- install dependency
```bash
glide get package_name
```

Initialize dev:
- install glide
- Replace all occurences of `golang-seed` by your project name.
- Replace all occurences of `github.com/dohr-michael` by your project path.
- install dependencies.
```bash
glide install
```
- run project :
```bash
go run main.go start
```
- run project with hot reload (http server)
```bash
go get github.com/codegangsta/gin
gin --appPort 8080 --buildArgs main.go -i run start
```

Docker
------

- Generate container :
```bash
make build
```

- Run container :
```bash
docker run -d [-v .my-config-file.yml:/.config.yml] project-name:version
```

TODO
----

- Build as CLI
- More documentation
- ci/cd, docker-hub integration ?
- Templating of the seed ?
  - GRPC
  - REST
  - ...

