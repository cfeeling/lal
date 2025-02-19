// Copyright 2019, Chef.  All rights reserved.
// https://github.com/cfeeling/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package httpflv

import (
	"testing"

	"github.com/q191201771/naza/pkg/assert"
)

func BenchmarkFlvFileReader(b *testing.B) {
	var tmp uint32
	for i := 0; i < b.N; i++ {
		var r FlvFileReader
		err := r.Open("testdata/test.flv")
		assert.Equal(b, nil, err)
		for {
			tag, err := r.ReadTag()
			if err != nil {
				break
			}
			tmp += uint32(tag.Raw[0])
		}
		r.Dispose()
	}
	//nazalog.Debug(tmp)
}

func BenchmarkCloneTag(b *testing.B) {
	var tmp uint32
	var r FlvFileReader
	err := r.Open("testdata/test.flv")
	assert.Equal(b, nil, err)
	tag, err := r.ReadTag()
	assert.Equal(b, nil, err)
	r.Dispose()
	for i := 0; i < b.N; i++ {
		tag2 := tag.clone()
		tmp += uint32(tag2.Raw[0])
	}
}
