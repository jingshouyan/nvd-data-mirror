package log

import (
	slog "log"
	"runtime"
)

func Println(v ...any) {
	showInfo()
	slog.Println(v...)
}

func Printf(format string, v ...any) {
	slog.Printf(format, v...)
}

func showInfo() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	slog.Printf("%d Kb\n", m.Alloc/1024)
}
