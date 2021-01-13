package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	rootCmd.Flags().BoolVarP(&absoluteTime, "absolute", "a", false, "Use absolute time for timestamps")
	rootCmd.Flags().StringVarP(&trimPrefix, "trim", "t", "", "Trim package prefix (ex: github.com/michaelquigley/)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])),
	Short: "pfxlog filter",
	Run:   Main,
}
var trimPrefix string
var absoluteTime bool

func Main(_ *cobra.Command, _ []string) {
	r := bufio.NewReader(os.Stdin)
	var last time.Time
	lastSet := false
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		msg := make(map[string]interface{})
		err = json.Unmarshal([]byte(line), &msg)
		if err != nil {
			fmt.Println(strings.TrimSpace(line))
			continue
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
		var prefix string
		if v, found := msg["func"]; found {
			prefix = strings.TrimPrefix(v.(string), trimPrefix)
		}
		if context, found := msg["context"]; found {
			prefix += " [" + context.(string) + "]"
		}
		message := msg["msg"].(string)
		data := data(msg)
		if len(data) > 0 {
			fields := "{"
			field := 0
			for k, v := range data {
				if field > 0 {
					fields += " "
				}
				field++
				fields += fmt.Sprintf("%s=[%v]", k, v)
			}
			fields += "} "
			message = ansi.LightCyan + fields + ansi.DefaultFG + message
		}
		var fmtTs string
		if absoluteTime {
			fmtTs = fmt.Sprintf("[%s]", last)
		} else {
			fmtTs = fmt.Sprintf("[%8.3f]", delta)
		}
		fmt.Printf("%s %s %s: %s\n",
			ansi.Blue+fmtTs+ansi.DefaultFG,
			level,
			ansi.Cyan+prefix+ansi.DefaultFG,
			message)
	}
}

func data(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range in {
		if k != "level" && k != "func" && k != "file" && k != "msg" && k != "time" && k != "context" {
			out[k] = v
		}
	}
	return out
}

var panicColor = ansi.Red + "  PANIC" + ansi.DefaultFG
var fatalColor = ansi.Red + "  FATAL" + ansi.DefaultFG
var errorColor = ansi.Red + "  ERROR" + ansi.DefaultFG
var warnColor = ansi.Yellow + "WARNING" + ansi.DefaultFG
var infoColor = ansi.White + "   INFO" + ansi.DefaultFG
var debugColor = ansi.Blue + "  DEBUG" + ansi.DefaultFG
var traceColor = ansi.LightBlack + "  TRACE" + ansi.DefaultFG
