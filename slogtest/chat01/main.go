package main

import (
	"log/slog"
	"time"
)

func main() {

	// slog初体验
	slog.Info("hello world", " time is now:", time.Now())

}
