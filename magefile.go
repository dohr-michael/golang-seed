// +build mage

package main

import "github.com/magefile/mage/sh"
import "github.com/magefile/mage/mg"

func init() {

}

func downloadDeps() error {
	return sh.Run("go", "mod", "download")
}

func generate() error {
	return sh.Run("go", "generate", "./...")
}

func Test() error {
	mg.Deps(downloadDeps, generate)
	return sh.Run("go", "test", "./...")
}

func Build() error {
	mg.Deps(Test)
	return sh.Run("go", "build", "./...")
}


func Snapshot() error {
	mg.Deps(Test)
	return sh.Run("goreleaser", "--rm-dist", "--snapshot")
}

func Release() error {
	mg.Deps(Test)
	return sh.Run("goreleaser", "--rm-dist")
}
