package main

import (
	_ "embed"

	"github.com/alveycoin/alveychain/command/root"
	"github.com/alveycoin/alveychain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
