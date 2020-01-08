// use test.sh

package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/clipperhouse/setwriter"
	"github.com/clipperhouse/typewriter"
)

func main() {
	// don't let bad test or gen'd files get us stuck
	filter := func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), "_test.go") && !strings.HasSuffix(f.Name(), "_set.go")
	}

	a, err := typewriter.NewAppFiltered("+test", filter)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := a.WriteAll(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
