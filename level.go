package logger

type Level int

var (
	levelStrings = [...]string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

func (l Level) String() string {
	if l < 0 || int(l) > len(levelStrings) {
		return "UNKNOWN"
	}
	return levelStrings[int(l)]
}
