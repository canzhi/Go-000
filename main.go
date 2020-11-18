package main

import (
	"os/exec"
)

func main() {
	//! git add .
	if err := exec.Command("git", "add", ".").Run(); err != nil {
		panic(err)
	}
	//! git commit -m commnet
	if err := exec.Command("git", "commit").Run(); err != nil {
		panic(err)
	}
	//! git push
	if err := exec.Command("git", "push").Run(); err != nil {
		panic(err)
	}
}
