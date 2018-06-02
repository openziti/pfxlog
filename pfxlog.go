package pfxlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"runtime"
	"path/filepath"
)

var trimmedPrefix = ""

func Global(level logrus.Level) {
	fmt := new(prefixed.TextFormatter)
	logrus.SetFormatter(fmt)
	logrus.SetLevel(level)
}

func SetTrimmedPrefix(prefix string) {
	trimmedPrefix = prefix
}

func Logger() *logrus.Entry {
	return logrus.StandardLogger().WithField("prefix", functionName())
}

func AttachedLogger(attached string) *logrus.Entry {
	return logrus.StandardLogger().WithField("prefix", fmt.Sprintf("%s//[%s]", functionName(), attached))
}

func functionName() string {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)

	funcName := f.Name()
	if trimmedPrefix != "" {
		if len(filepath.SplitList(funcName)) > 0 {
			trimmedName, err := filepath.Rel(trimmedPrefix, funcName)
			if err != nil {
				return funcName + "(!)"
			} else {
				return trimmedName
			}
		}
	}
	return funcName
}
