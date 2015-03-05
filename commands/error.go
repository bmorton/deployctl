package commands

import (
	"fmt"
	"os"
)

func errorAndBail(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}
