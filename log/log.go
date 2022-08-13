package log

import (
	slog "log"
)

func Println(v ...any) {
	slog.Println(v...)

}

func Printf(format string, v ...any) {
	slog.Printf(format, v...)
}
