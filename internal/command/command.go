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
	input = flag.String("s", "", "input file")
	line  = flag.Int("l", 0, "line number containing the target function")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: goffn [-l line number] [-s input file]\n")
	flag.PrintDefaults()
}

func exec() int {
	flag.Usage = usage
	flag.Parse()

	if len(*input) == 0 {
		fmt.Fprintln(os.Stderr, "must specify filename")
		return 1
	}

	anal := analysis.New(
		*input,
		*line,
	)

	if err := anal.Do(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute : %s\n", err.Error())
		return 1
	}

	return 0
}
