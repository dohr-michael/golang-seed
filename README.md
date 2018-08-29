golang-seed
===========

Golang seed application with yaml config file + cli

Default command provided:
- start : to start the process.
- version : to display the version (based to git commit / branch / build time)

Package management by [https://github.com/Masterminds/glide](Glide)
- install dependency
```bash
glide get package_name
```

Initialize dev:
- install glide
- Replace all occurences of `golang-seed` by your project name.
- install dependencies.
```bash
glide install
```
- run project :
```bash
go run main.go start
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


