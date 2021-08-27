package color

import "github.com/fatih/color"

type Color int

const (
	Cyan Color = iota
	Blue
	Red
	Green
	Yellow
	Magenta
	Gray
)

func (c Color) new() *color.Color {
	return [...]*color.Color{
		color.New(color.FgCyan),
		color.New(color.FgBlue),
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgMagenta),
		color.New(color.Italic),
	}[c]
}

func (c Color) Println(args ...interface{}) {
	c.new().Println(args...)
}

func (c Color) Printf(format string, args ...interface{}) {
	c.new().Printf(format, args...)
}

func (c Color) Sprintf(format string, args ...interface{}) string {
	return c.new().Sprintf(format, args...)
}
