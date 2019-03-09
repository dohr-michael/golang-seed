// +build mage

package main

import "github.com/magefile/mage/sh"
import "github.com/magefile/mage/mg"

func init() {

}

func downloadDeps() error {
	return sh.RunV("go", "mod", "download")
}

func generate() error {
	return sh.RunV("go", "generate", "./...")
}

func Run() error {
	mg.Deps(downloadDeps, generate)
	return sh.RunV("gin", "--appPort", "8080", "--buildArgs", "main.go", "-i", "run", "start")
}

func Test() error {
	mg.Deps(downloadDeps, generate)
	return sh.RunV("go", "test", "./...")
}

func Build() error {
	mg.Deps(Test)
	return sh.RunV("go", "build", "./...")
}


func Snapshot() error {
	mg.Deps(Test)
	return sh.RunV("goreleaser", "--rm-dist", "--snapshot")
}

func Release() error {
	mg.Deps(Test)
	return sh.RunV("goreleaser", "--rm-dist")
}
