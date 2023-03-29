package dbutil

import (
	"gorm.io/gorm/logger"
	"strings"
)

func ParseLevel(lvl string) logger.LogLevel {
	switch strings.ToLower(lvl) {
	case "info":
		return logger.Info
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	default:
		return logger.Silent
	}
}
