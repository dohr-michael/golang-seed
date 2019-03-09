golang-Seed
===========

[![Build Status](https://travis-ci.org/dohr-michael/golang-seed.svg?branch=master)](https://travis-ci.org/dohr-michael/golang-seed)


- Using [go mod](https://github.com/golang/go/wiki/Modules)
- Build in with [magefile](https://magefile.org/)
- Release management with [GoReleaser](https://goreleaser.com/)

- Needs
    - Replacement in the project (in order):
        - `golang-seed` by your project name
        - `github.com/dohr-michael` by your github / git path / by your go package root.
        - `jdrbahamut` by your docker registry organization.
    - environment variables
        - GITHUB_TOKEN
    - Docker registry authentication
- Release Generate :
    - Binaries
        - darwin_amd64
        - darwin_386
        - linux_amd64
        - linux_arm64
        - linux_386
        - windows_amd64
        - windows_386
    - Docker Container
        - linux_amd64
    - Github release
    - Homebrew Tap
    - Scoop recipe