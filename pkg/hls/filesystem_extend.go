package hls

import "github.com/q191201771/naza/pkg/filesystemlayer"

func GetFsHandler() filesystemlayer.IFileSystemLayer {
	return fslCtx
}
