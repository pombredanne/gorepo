// main.go
package main

import (
	"fmt"

	rpm "github.com/cavaliercoder/go-rpm"
)

func main() {
	p, err := rpm.OpenPackageFile("../testdata/python-colorama-0.3.2-3.el7.noarch.rpm")
	if err != nil {
		fmt.Errorf("error reading %s: %v\n", err)
	}
	fmt.Printf("%v", (*p).Epoch())
}
