// Copyright 2021, Chef.  All rights reserved.
// https://github.com/cfeeling/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package hls

import (
	"sync"

	"github.com/cfeeling/naza/pkg/filesystemlayer"
)

var (
	fslCtx  filesystemlayer.IFileSystemLayer
	setOnce sync.Once
)

func SetUseMemoryAsDiskFlag(flag bool) {
	setOnce.Do(func() {
		var t filesystemlayer.FSLType
		if flag {
			t = filesystemlayer.FSLTypeMemory
		} else {
			t = filesystemlayer.FSLTypeDisk
		}
		if fslCtx == nil || fslCtx.Type() != t {
			fslCtx = filesystemlayer.FSLFactory(t)
		}
	})
}

func RemoveAll(path string) error {
	return fslCtx.RemoveAll(path)
}

func init() {
	fslCtx = filesystemlayer.FSLFactory(filesystemlayer.FSLTypeDisk)
}
