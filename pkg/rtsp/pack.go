// Copyright 2020, Chef.  All rights reserved.
// https://github.com/cfeeling/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtsp

import (
	"fmt"
	"time"

	"github.com/cfeeling/lal/pkg/base"
)

// rfc2326 10.1 OPTIONS

// CSeq
var ResponseOptionsTmpl = "RTSP/1.0 200 OK\r\n" +
	"Server: " + base.LalRtspOptionsResponseServer + "\r\n" +
	"CSeq: %s\r\n" +
	"Public: DESCRIBE, ANNOUNCE, SETUP, PLAY, PAUSE, RECORD, TEARDOWN\r\n" +
	"\r\n"

// rfc2326 10.3 ANNOUNCE
//var RequestAnnounceTmpl = "not impl"

// CSeq
var ResponseAnnounceTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"\r\n"

// rfc2326 10.2 DESCRIBE

// CSeq, Date, Content-Length,
var ResponseDescribeTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Date: %s\r\n" +
	"Content-Type: application/sdp\r\n" +
	"Content-Length: %d\r\n" +
	"\r\n" +
	"%s"

// rfc2326 10.4 SETUP
// CSeq, Date, Session, Transport
var ResponseSetupTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Date: %s\r\n" +
	"Session: %s\r\n" +
	"Transport: %s\r\n" +
	"\r\n"

// rfc2326 10.11 RECORD
//var RequestRecordTmpl = "not impl"

// CSeq, Session
var ResponseRecordTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Session: %s\r\n" +
	"\r\n"

// rfc2326 10.5 PLAY

// CSeq Date
var ResponsePlayTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"Date: %s\r\n" +
	"\r\n"

// rfc2326 10.7 TEARDOWN
//var RequestTeardownTmpl = "not impl"

// CSeq
var ResponseTeardownTmpl = "RTSP/1.0 200 OK\r\n" +
	"CSeq: %s\r\n" +
	"\r\n"

func PackResponseOptions(cseq string) string {
	return fmt.Sprintf(ResponseOptionsTmpl, cseq)
}

func PackResponseAnnounce(cseq string) string {
	return fmt.Sprintf(ResponseAnnounceTmpl, cseq)
}

func PackResponseDescribe(cseq, sdp string) string {
	date := time.Now().Format(time.RFC1123)
	return fmt.Sprintf(ResponseDescribeTmpl, cseq, date, len(sdp), sdp)
}

func PackResponseSetup(cseq string, htv string) string {
	date := time.Now().Format(time.RFC1123)

	return fmt.Sprintf(ResponseSetupTmpl, cseq, date, sessionId, htv)
}

func PackResponseRecord(cseq string) string {
	return fmt.Sprintf(ResponseRecordTmpl, cseq, sessionId)
}

func PackResponsePlay(cseq string) string {
	date := time.Now().Format(time.RFC1123)
	return fmt.Sprintf(ResponsePlayTmpl, cseq, date)
}

func PackResponseTeardown(cseq string) string {
	return fmt.Sprintf(ResponseTeardownTmpl, cseq)
}

// body 可以为空
func PackRequest(method, uri string, headers map[string]string, body string) (ret string) {
	ret = method + " " + uri + " RTSP/1.0\r\n"
	for k, v := range headers {
		ret += k + ": " + v + "\r\n"
	}
	ret += "\r\n"

	if body != "" {
		ret += body
	}

	return ret
}
