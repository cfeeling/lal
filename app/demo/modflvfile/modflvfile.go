// Copyright 2019, Chef.  All rights reserved.
// https://github.com/cfeeling/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"flag"
	"io"
	"time"

	"github.com/cfeeling/lal/pkg/base"

	"github.com/cfeeling/lal/pkg/httpflv"
	log "github.com/q191201771/naza/pkg/nazalog"
)

// 修改flv文件的一些信息（比如某些tag的时间戳）后另存文件
//
// Usage:
// ./bin/modflvfile -i /tmp/in.flv -o /tmp/out.flv

var countA int
var countV int
var exitFlag bool

func hookTag(tag *httpflv.Tag) {
	log.Infof("%+v", tag.Header)
	if tag.Header.Timestamp != 0 {
		tag.ModTagTimestamp(tag.Header.Timestamp + uint32(time.Now().Unix()/1e6))
	}
}

func main() {
	var err error
	inFileName, outFileName := parseFlag()

	var ffr httpflv.FLVFileReader
	err = ffr.Open(inFileName)
	log.Assert(nil, err)
	defer ffr.Dispose()
	log.Infof("open input flv file succ.")

	var ffw httpflv.FLVFileWriter
	err = ffw.Open(outFileName)
	log.Assert(nil, err)
	defer ffw.Dispose()
	log.Infof("open output flv file succ.")

	flvHeader, err := ffr.ReadFLVHeader()
	log.Assert(nil, err)

	err = ffw.WriteRaw(flvHeader)
	log.Assert(nil, err)

	for {
		tag, err := ffr.ReadTag()
		if err == io.EOF {
			log.Infof("EOF.")
			break
		}
		log.Assert(nil, err)
		hookTag(&tag)
		err = ffw.WriteRaw(tag.Raw)
		log.Assert(nil, err)
	}
}

func parseFlag() (string, string) {
	i := flag.String("i", "", "specify input flv file")
	o := flag.String("o", "", "specify ouput flv file")
	flag.Parse()
	if *i == "" || *o == "" {
		flag.Usage()
		base.OSExitAndWaitPressIfWindows(1)
	}
	return *i, *o
}
