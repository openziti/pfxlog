package main

import (
	"encoding/json"
	"fmt"
	"github.com/mgutz/ansi"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	trim := ""
	if len(os.Args) > 1 {
		trim = os.Args[1]
	}
	fmt.Printf("trimming: [%s]\n", trim)
	decoder := json.NewDecoder(os.Stdin)
	var last time.Time
	lastSet := false
	for {
		msg := make(map[string]interface{})
		err := decoder.Decode(&msg)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		stamp, err := time.Parse(time.RFC3339, msg["time"].(string))
		if err != nil {
			panic(err)
		}
		if !lastSet {
			last = stamp
			lastSet = true
		}
		delta := stamp.Sub(last).Seconds()
		var level string
		switch msg["level"].(string) {
		case "panic":
			level = panicColor
		case "fatal":
			level = fatalColor
		case "error":
			level = errorColor
		case "warning":
			level = warnColor
		case "info":
			level = infoColor
		case "debug":
			level = debugColor
		case "trace":
			level = traceColor
		default:
			panic(fmt.Errorf("unknown (%s)", msg["level"].(string)))
		}
		prefix := strings.TrimPrefix(msg["func"].(string), trim)
		if context, found := msg["context"]; found {
			prefix += " [" + context.(string) + "]"
		}
		fmt.Printf("%s %s %s: %s\n",
			ansi.Blue+fmt.Sprintf("[%8.3f]", delta)+ansi.DefaultFG,
			level,
			ansi.Cyan+prefix+ansi.DefaultFG,
			msg["msg"])
	}
}

var panicColor = ansi.Red + "  PANIC" + ansi.DefaultFG
var fatalColor = ansi.Red + "  FATAL" + ansi.DefaultFG
var errorColor = ansi.Red + "  ERROR" + ansi.DefaultFG
var warnColor = ansi.Yellow + "WARNING" + ansi.DefaultFG
var infoColor = ansi.White + "   INFO" + ansi.DefaultFG
var debugColor = ansi.Blue + "  DEBUG" + ansi.DefaultFG
var traceColor = ansi.LightBlack + "  TRACE" + ansi.DefaultFG
