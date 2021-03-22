// Copyright 2020, Chef.  All rights reserved.
// https://github.com/cfeeling/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package logic

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/cfeeling/lal/pkg/base"

	"github.com/cfeeling/naza/pkg/bininfo"
	"github.com/cfeeling/naza/pkg/nazalog"
	//"github.com/felixge/fgprof"
)

var (
	config *Config
	sm     *ServerManager
)

func Entry(confFile string) {
	configTmp := loadConf(confFile)
	RunServer(configTmp)
}

func RunServer(c *Config) {
	config = c
	initLog(c.LogConfig)
	nazalog.Infof("bininfo: %s", bininfo.StringifySingleLine())
	nazalog.Infof("version: %s", base.LALFullInfo)
	nazalog.Infof("github: %s", base.LALGithubSite)
	nazalog.Infof("doc: %s", base.LALDocSite)

	sm = NewServerManager()

	if c.PProfConfig.Enable {
		go runWebPProf(c.PProfConfig.Addr)
	}
	go runSignalHandler(func() {
		sm.Dispose()
	})

	sm.RunLoop()
}

func Dispose() {
	sm.Dispose()
}

func loadConf(confFile string) *Config {
	config, err := LoadConf(confFile)
	if err != nil {
		nazalog.Errorf("load conf failed. file=%s err=%+v", confFile, err)
		os.Exit(1)
	}
	nazalog.Infof("load conf file succ. file=%s content=%+v", confFile, config)
	return config
}

func initLog(opt nazalog.Option) {
	if err := nazalog.Init(func(option *nazalog.Option) {
		*option = opt
	}); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "initial log failed. err=%+v\n", err)
		os.Exit(1)
	}
	nazalog.Info("initial log succ.")
}

func runWebPProf(addr string) {
	nazalog.Infof("start web pprof listen. addr=%s", addr)

	//http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())

	if err := http.ListenAndServe(addr, nil); err != nil {
		nazalog.Error(err)
		return
	}
}
