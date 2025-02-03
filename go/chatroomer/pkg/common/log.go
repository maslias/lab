package common

import "go.uber.org/zap"

type HttpLog struct {
	*zap.SugaredLogger
}

func NewHttpLog() *HttpLog {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	return &HttpLog{logger}
}
