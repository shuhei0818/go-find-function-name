package command

import (
	"flag"
	"fmt"
	"os"

	"github.com/shuhei0818/go-find-function-name/internal/analysis"
)

func Exec() {
	os.Exit(exec())
}

var (
	line        = flag.Int("l", 0, "line number containing the target function")
	showVersion = flag.Bool("v", false, "Print the version")
)

var (
	name    = "goffn"
	version = "0.2.0"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: goffn [-l line_number] filename\n")
	flag.PrintDefaults()
}

func exec() int {
	flag.Usage = usage
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s v%s\n", name, version)
		return 0
	}

	input := flag.Arg(0)
	if input == "" {
		fmt.Fprintln(os.Stderr, "must specify filename")
		flag.Usage()
		return 1
	}

	anal := analysis.New(
		input,
		*line,
	)

	if err := anal.Do(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute : %s\n", err.Error())
		return 1
	}

	return 0
}
