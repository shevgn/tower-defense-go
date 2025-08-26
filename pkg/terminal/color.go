package terminal

// Color renderer to colorize text
type Color string

const (
	ColorReset   Color = "\033[0m"
	ColorBlack   Color = "\033[30m"
	ColorRed     Color = "\033[31m"
	ColorGreen   Color = "\033[32m"
	ColorYellow  Color = "\033[33m"
	ColorBlue    Color = "\033[34m"
	ColorMagenta Color = "\033[35m"
	ColorCyan    Color = "\033[36m"
	ColorWhite   Color = "\033[37m"

	ColorBoldBlack   Color = "\033[1;30m"
	ColorBoldRed     Color = "\033[1;31m"
	ColorBoldGreen   Color = "\033[1;32m"
	ColorBoldYellow  Color = "\033[1;33m"
	ColorBoldBlue    Color = "\033[1;34m"
	ColorBoldMagenta Color = "\033[1;35m"
	ColorBoldCyan    Color = "\033[1;36m"
	ColorBoldWhite   Color = "\033[1;37m"
)
