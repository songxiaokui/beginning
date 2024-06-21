package daemon

import "context"

type daemonConfig struct {
	confConfig
	logConfig
}
type confConfig struct {
	path string
}
type logConfig struct {
	useConsole bool
	path       string
	maxSize    int
	maxAge     int
	maxBackups int
}

type DaemonProcess struct {
	ctx context.Context // 控制整个进程的上下文
	dc  daemonConfig    // 进程配置
}
