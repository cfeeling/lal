package hls

type ReadFileFallback func(rootOutPath string, fileName string, streamName string, fileType string) ([]byte, error)
