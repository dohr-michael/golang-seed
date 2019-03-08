// +build mage

package main

import "github.com/magefile/mage/sh"
import "github.com/magefile/mage/mg"

func init() {

}

func Test() error {
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
