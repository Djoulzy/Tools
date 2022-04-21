package confload

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

func tryingFile(confFile string) error {
	if _, err := os.Stat(confFile); os.IsNotExist(err) {
		log.Printf("Trying to load conf file %s ... err", confFile)
		return errors.New("can't find conf file")
	}
	log.Printf("Trying to load conf file %s ... OK", confFile)
	return nil
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
			if tryingFile(confFile) != nil {
				confFile = fmt.Sprintf("./%s", iniName)
			} else {
				found = true
			}
		} else {
			found = true
		}
	}

	if !found {
		if tryingFile(confFile) != nil {
			log.Printf("No conf file found, using default values")
			return errors.New("can't find conf file")
		}
	}

	log.Printf("Loading Conf File ... %s", confFile)
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
