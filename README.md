golang-seed
===========

Golang seed application with yaml config file + cli

Default command provided:
- start : to start the process.
- version : to display the version (based to git commit / branch / build time)

Package management by [https://github.com/Masterminds/glide](Glide)
- install dependency
```shell
glide get package_name
```

Docker
------

- Generate container :
```shell
make build
```

- Run container :
```shell
docker run -d [-v .my-config-file.yml:/.config.yml] project-name:version
```


