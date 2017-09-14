package clog

import (
	"fmt"
	"log"
)

var LogLevel int
var StartLogging bool

var fg_colors = map[string]string{
	"black":        "0;30",
	"dark_gray":    "1;30",
	"blue":         "0;34",
	"light_blue":   "1;34",
	"green":        "0;32",
	"light_green":  "1;32",
	"cyan":         "0;36",
	"light_cyan":   "1;36",
	"red":          "0;31",
	"light_red":    "1;31",
	"purple":       "0;35",
	"light_purple": "1;35",
	"brown":        "0;33",
	"yellow":       "1;33",
	"light_gray":   "0;37",
	"white":        "1;37",
}

var bg_colors = map[string]string{
	"black":      "40",
	"red":        "41",
	"green":      "42",
	"yellow":     "43",
	"blue":       "44",
	"magenta":    "45",
	"cyan":       "46",
	"light_gray": "47",
}

type errorsColors struct {
	name  string
	level int
	fg    string
	bg    string
}

var TEST = errorsColors{
	name:  "TEST",
	level: 0,
	fg:    "green",
	bg:    "black",
}
var DEBUG = errorsColors{
	name:  "DBUG",
	level: 5,
	fg:    "dark_gray",
	bg:    "black",
}
var INFO = errorsColors{
	name:  "INFO",
	level: 4,
	fg:    "light_gray",
	bg:    "black",
}
var TRACE = errorsColors{
	name:  "TRAC",
	level: 4,
	fg:    "white",
	bg:    "black",
}
var WARN = errorsColors{
	name:  "WARN",
	level: 2,
	fg:    "yellow",
	bg:    "black",
}
var ERROR = errorsColors{
	name:  "ERRR",
	level: 1,
	fg:    "light_red",
	bg:    "black",
}
var FATAL = errorsColors{
	name:  "FATAL",
	level: 1,
	fg:    "white",
	bg:    "red",
}

func getColoredString(str string, fgcolor string, bgcolor string) string {
	colored_string := ""

	if len(fgcolor) != 0 {
		if len(fg_colors[fgcolor]) != 0 {
			colored_string = fmt.Sprintf("%s%c[%sm", colored_string, 27, fg_colors[fgcolor])
		}
	}

	if len(bgcolor) != 0 {
		if len(bg_colors[bgcolor]) != 0 {
			colored_string = fmt.Sprintf("%s%c[%sm", colored_string, 27, bg_colors[bgcolor])
		}
	}

	colored_string = fmt.Sprintf("%s%s%c[0m", colored_string, str, 27)
	// return $colored_string;
	return colored_string
}

func Println(fgcolor string, bgcolor string, str string) {
	tmp := getColoredString(str, fgcolor, bgcolor)
	log.Println(tmp)
}

func Printf(fgcolor string, bgcolor string, format string, vars ...interface{}) {
	tmp := getColoredString(format, fgcolor, bgcolor)
	log.Printf(tmp, vars...)
}

func Output(str string, vars ...interface{}) {
	before := fmt.Sprintf("%s", str)
	Printf(INFO.fg, INFO.bg, before, vars...)
}

func logOutput(etype errorsColors, pack string, function string, str string, vars ...interface{}) {
	if LogLevel < etype.level || StartLogging == false {
		return
	}
	before := fmt.Sprintf("%s|%s|%s| %s", etype.name, pack, function, str)
	Printf(etype.fg, etype.bg, before, vars...)
}

func Warn(pack string, function string, str string, vars ...interface{}) {
	logOutput(WARN, pack, function, str, vars...)
}

func Info(pack string, function string, str string, vars ...interface{}) {
	logOutput(INFO, pack, function, str, vars...)
}

func Debug(pack string, function string, str string, vars ...interface{}) {
	logOutput(DEBUG, pack, function, str, vars...)
}

func Test(pack string, function string, str string, vars ...interface{}) {
	logOutput(TEST, pack, function, str, vars...)
}

func Error(pack string, function string, str string, vars ...interface{}) {
	logOutput(ERROR, pack, function, str, vars...)
}

func Fatal(pack string, function string, err error) {
	logOutput(FATAL, pack, function, "%s", err)
	log.Fatal()
}

func Trace(pack string, function string, str string, vars ...interface{}) {
	logOutput(TRACE, pack, function, str, vars...)
}
