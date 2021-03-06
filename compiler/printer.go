// +build !windows

package compiler

import (
	"os"

	"github.com/alecthomas/chroma/quick"
)

func PrettyPrintSource(source string) {
	quick.Highlight(os.Stdout, source, "go", "terminal", "vim")
}
