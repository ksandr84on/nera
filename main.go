package main

import (
	_ "embed"

	"github.com/ksandr84on/nera/command/root"
	"github.com/ksandr84on/nera/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
