package pfxlog

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
	"strings"
)

type Formatter struct{}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s [%s] %s: %s\n",
		ansi.Blue + entry.Time.String() + ansi.DefaultFG,
		ansi.Green + strings.TrimPrefix(entry.Caller.Function, prefix) + ansi.DefaultFG,
		ansi.Yellow + entry.Level.String() + ansi.DefaultFG, entry.Message),
	),
	nil
}
