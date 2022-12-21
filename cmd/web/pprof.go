package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/grin-ch/dever-box-api/cfg"
)

func init() {
	if cfg.Config.Pprof.Enable {
		go func() {
			http.ListenAndServe(fmt.Sprintf(":%d", cfg.Config.Pprof.Port), nil)
		}()
	}

}
