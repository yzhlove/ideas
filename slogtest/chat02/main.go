package main

import (
	"log/slog"
	"os"
)

func main() {

	// slog结构化输出日志

	// TextHandle
	handle := slog.NewTextHandler(os.Stderr, nil)
	l := slog.New(handle)
	l.Info("this is info", "level")
	l.Warn("this is warn", "level")
	l.Error("this is error", "level")

	// JsonHandle
	handle2 := slog.NewJSONHandler(os.Stderr, nil)
	l = slog.New(handle2)
	l.Info("this is info", "level")
	l.Warn("this is warn", "level")
	l.Error("this is error", "level")

}
