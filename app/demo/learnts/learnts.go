// Copyright 2020, Chef.  All rights reserved.
// https://github.com/cfeeling/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package main

import (
	"io/ioutil"

	"github.com/cfeeling/lal/pkg/mpegts"

	"github.com/cfeeling/lal/pkg/hls"
	"github.com/cfeeling/naza/pkg/nazalog"
)

// 学习如何解析TS文件。注意，该程序还没有写完。

var (
	pat        mpegts.PAT
	pmt        mpegts.PMT
	pid2stream map[uint16]*Stream
)

type Stream struct {
}

var filename = "/Volumes/Data/nrm-0.ts"

func handlePacket(packet []byte) {
	h := mpegts.ParseTSPacketHeader(packet)
	index := 4
	nazalog.Debugf("%+v", h)

	var adaptation mpegts.TSPacketAdaptation
	switch h.Adaptation {
	case mpegts.AdaptationFieldControlNo:
		// noop
	case mpegts.AdaptationFieldControlFollowed:
		adaptation = mpegts.ParseTSPacketAdaptation(packet[4:])
		index++
	default:
		nazalog.Warn(h.Adaptation)
	}
	index += int(adaptation.Length)

	if h.Pid == mpegts.PidPAT {
		if h.PayloadUnitStart == 1 {
			index++
		}
		pat = mpegts.ParsePAT(packet[index:])
		nazalog.Debugf("%+v", pat)
		return
	}

	if pat.SearchPID(h.Pid) {
		if h.PayloadUnitStart == 1 {
			index++
		}
		pmt = mpegts.ParsePMT(packet[index:])
		nazalog.Debugf("%+v", pmt)

		for _, ele := range pmt.ProgramElements {
			pid2stream[ele.Pid] = &Stream{}
		}
		return
	}

	_, ok := pid2stream[h.Pid]
	if !ok {
		nazalog.Warn(h.Pid)
	}

	// 判断是否有PES
	if h.PayloadUnitStart == 1 {
		pes, length := mpegts.ParsePES(packet[index:])
		nazalog.Debugf("%+v, %d", pes, length)
	}
}

func main() {
	_ = nazalog.Init(func(option *nazalog.Option) {
		option.AssertBehavior = nazalog.AssertFatal
	})
	defer nazalog.Sync()

	pid2stream = make(map[uint16]*Stream)

	content, err := ioutil.ReadFile(filename)
	nazalog.Assert(nil, err)

	packets, _ := hls.SplitFragment2TSPackets(content)

	for _, packet := range packets {
		handlePacket(packet)
	}
}
