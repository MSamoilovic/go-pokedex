package main

const (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
	colorBold    = "\033[1m"
)

func colorize(color, text string) string {
	return color + text + colorReset
}

var typeColors = map[string]string{
	"normal":   colorWhite,
	"fire":     colorRed,
	"water":    colorBlue,
	"grass":    colorGreen,
	"electric": colorYellow,
	"ice":      colorCyan,
	"fighting": colorRed,
	"poison":   colorMagenta,
	"ground":   colorYellow,
	"flying":   colorCyan,
	"psychic":  colorMagenta,
	"bug":      colorGreen,
	"rock":     colorYellow,
	"ghost":    colorMagenta,
	"dragon":   colorBlue,
	"dark":     colorWhite,
	"steel":    colorCyan,
	"fairy":    colorMagenta,
}

func typeColor(typeName string) string {
	if color, ok := typeColors[typeName]; ok {
		return colorize(color, typeName)
	}
	return typeName
}
