package terminal

// Color renderer to colorize text
type Color string

const (
	ColorReset Color = "\033[0m"
	FgBlack    Color = "\033[30m"
	FgRed      Color = "\033[31m"
	FgGreen    Color = "\033[32m"
	FgYellow   Color = "\033[33m"
	FgBlue     Color = "\033[34m"
	FgMagenta  Color = "\033[35m"
	FgCyan     Color = "\033[36m"
	FgWhite    Color = "\033[37m"

	FgBoldBlack   Color = "\033[1;30m"
	FgBoldRed     Color = "\033[1;31m"
	FgBoldGreen   Color = "\033[1;32m"
	FgBoldYellow  Color = "\033[1;33m"
	FgBoldBlue    Color = "\033[1;34m"
	FgBoldMagenta Color = "\033[1;35m"
	FgBoldCyan    Color = "\033[1;36m"
	FgBoldWhite   Color = "\033[1;37m"

	BgBlack   Color = "\033[40m"
	BgRed     Color = "\033[41m"
	BgGreen   Color = "\033[42m"
	BgYellow  Color = "\033[43m"
	BgBlue    Color = "\033[44m"
	BgMagenta Color = "\033[45m"
	BgCyan    Color = "\033[46m"
	BgWhite   Color = "\033[47m"

	BgBoldBlack   Color = "\033[1;40m"
	BgBoldRed     Color = "\033[1;41m"
	BgBoldGreen   Color = "\033[1;42m"
	BgBoldYellow  Color = "\033[1;43m"
	BgBoldBlue    Color = "\033[1;44m"
	BgBoldMagenta Color = "\033[1;45m"
	BgBoldCyan    Color = "\033[1;46m"
	BgBoldWhite   Color = "\033[1;47m"
)
