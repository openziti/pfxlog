package pfxlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Formatter struct{}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%s] %s: %s\n", entry.Caller.Function, entry.Level.String(), entry.Message)), nil
}
