package recover

import (
	"log/slog"
)

func Recover() {
	if err := recover(); err != nil {
		slog.Error("Recovered from panic:", err)
	}
}
