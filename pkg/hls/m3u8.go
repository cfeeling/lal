// Copyright 2020, Chef.  All rights reserved.
// https://github.com/cfeeling/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"bytes"
	"fmt"
	"strconv"
)

// content     需写入文件的内容
// filename    m3u8文件名
// filenameBak m3u8临时文件名
//
func writeM3u8File(content []byte, filename string, filenameBak string) error {
	if err := fslCtx.WriteFile(filenameBak, content, 0666); err != nil {
		return err
	}

	return fslCtx.Rename(filenameBak, filename)
}

// 如果当前duration比原m3u8文件的`EXT-X-TARGETDURATION`大，则更新`EXT-X-TARGETDURATION`的值
//
// content      原m3u8文件的内容
// currDuration 当前duration
//
// @return 处理后的m3u8文件内容
func updateTargetDurationInM3u8(content []byte, currDuration int) ([]byte, error) {
	l := bytes.Index(content, []byte("#EXT-X-TARGETDURATION:"))
	if l == -1 {
		return content, ErrHls
	}
	r := bytes.Index(content[l:], []byte{'\n'})
	if r == -1 {
		return content, ErrHls
	}
	oldDurationStr := bytes.TrimPrefix(content[l:l+r], []byte("#EXT-X-TARGETDURATION:"))
	oldDuration, err := strconv.Atoi(string(oldDurationStr))
	if err != nil {
		return content, err
	}
	if currDuration > oldDuration {
		tmpContent := make([]byte, l)
		copy(tmpContent, content[:l])
		tmpContent = append(tmpContent, []byte(fmt.Sprintf("#EXT-X-TARGETDURATION:%d", currDuration))...)
		tmpContent = append(tmpContent, content[l+r:]...)
		content = tmpContent
	}
	return content, nil
}

// content 传入m3u8文件内容
//
// @return durationSec m3u8中所有ts的时间总和。注意，使用的是m3u8文件中描述的ts时间，而不是读取ts文件中实际音视频数据的时间。
//
func CalcM3u8Duration(content []byte) (durationSec float64, err error) {
	lines := bytes.Split(content, []byte{'\n'})
	for _, line := range lines {
		if bytes.HasPrefix(line, []byte("#EXTINF:")) {
			line = bytes.TrimSpace(line)
			v := bytes.TrimSuffix(bytes.TrimPrefix(line, []byte("#EXTINF:")), []byte{','})
			v = bytes.TrimSpace(v)
			vv, err := strconv.ParseFloat(string(v), 64)
			if err != nil {
				return durationSec, err
			}
			durationSec += vv
		}
	}
	return
}
