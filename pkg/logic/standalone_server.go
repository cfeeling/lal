package logic

import (
	"github.com/cfeeling/lal/pkg/base"
	"github.com/cfeeling/lal/pkg/hls"
	"github.com/cfeeling/lal/pkg/rtmp"
	"github.com/cfeeling/lal/pkg/rtsp"
)

func NewStandaloneServerManager(config *Config) *ServerManager {
	m := &ServerManager{
		groupMap: make(map[string]*Group),
		exitChan: make(chan struct{}),
	}

	if config.HttpflvConfig.Enable || config.HttpflvConfig.EnableHttps ||
		config.HttptsConfig.Enable || config.HttptsConfig.EnableHttps ||
		config.HlsConfig.Enable || config.HlsConfig.EnableHttps {
		m.httpServerManager = base.NewHttpServerManager()
		m.httpServerHandler = NewHttpServerHandler(m)
		m.hlsServerHandler = hls.NewServerHandler(config.HlsConfig.OutPath)
	}

	if config.HlsConfig.Enable && config.HlsConfig.HttpReadFileFallback != nil {
		m.hlsServerHandler.SetReadFileFallback(config.HlsConfig.HttpReadFileFallback)
	}

	if config.RtmpConfig.Enable {
		m.rtmpServer = rtmp.NewServer(m, config.RtmpConfig.Addr)
	}
	if config.RtspConfig.Enable {
		m.rtspServer = rtsp.NewServer(config.RtspConfig.Addr, m)
	}
	if config.HttpApiConfig.Enable {
		m.httpApiServer = NewHttpApiServer(config.HttpApiConfig.Addr, m)
	}
	return m
}
