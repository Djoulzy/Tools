package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Djoulzy/Tools/clog"
	"github.com/go-ini/ini"
)

func tryingFile(confFile string) error {
	clog.Output("Trying to load conf file %s ...", confFile)
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		return errors.New("Can't find conf file")
	} else {
		return nil
	}
}

// Load de conf ini file and initialize the struct
func Load(iniName string, data interface{}) error {
	var confFile string
	var found bool = false

	flag.StringVar(&confFile, "f", "", "Config file path")
	flag.Parse()

	if len(confFile) == 0 {
		gopath := os.Getenv("GOPATH")
		if len(gopath) == 0 {
			gopath = ""
		}
		etcpath := strings.TrimRight(gopath, "/")
		confFile = fmt.Sprintf("%s/etc/%s", etcpath, iniName)
		if tryingFile(confFile) != nil {
			confFile = fmt.Sprintf("/etc/%s", iniName)
		} else {
			found = true
		}
	}

	if !found {
		if tryingFile(confFile) != nil {
			clog.Output("No conf file found, using default values")
			return errors.New("Can't find conf file")
		}
	}

	clog.Output("Loading Conf File ... %s", confFile)
	cfg, err := ini.Load(confFile)
	if err != nil {
		return err
	}

	err = cfg.MapTo(data)
	if err != nil {
		return err
	}
	return nil
}
