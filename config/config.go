package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/Djoulzy/Tools/clog"
	"github.com/go-ini/ini"
)

func Load(iniName string, data interface{}) interface{} {
	etcpath := strings.TrimRight(os.Getenv("GOPATH"), "/")
	confFile := fmt.Sprintf("%s/etc/%s", etcpath, iniName)
	clog.Output("Loading Conf File ... %s", confFile)

	cfg, err := ini.Load(confFile)
	if err != nil {
		clog.Fatal("config", "Load", err)
	}
	err = cfg.MapTo(data)
	return data
}
