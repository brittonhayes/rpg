package logger

import (
	"fmt"
	"io"
	"text/template"

	"github.com/fatih/color"
)

type Logger struct {
	Writer io.Writer
}

type Level int

const (
	LevelInfo Level = iota
	LevelChat
	LevelCombat
	LevelError
)

const (
	templateAttack = `{{.Attacker}} attacked {{.Target}} with {{.Attack}} for {{.Damage}} pts of damage`
	templateStatus = `Name={{.Character}} Health={{.Health}} Armor={{.Armor}}`
)

type Attack interface {
	HasName() string
	HasDamage() float64
}

func NewLogger(writer io.Writer) *Logger {
	return &Logger{Writer: writer}
}

func (l Level) String() string {
	return [...]string{"INFO", "CHAT", "COMBAT", "ERROR"}[l]
}

func selectPrefix(level Level) string {
	var prefix string
	msgFmt := "%-15s "
	switch level {
	case LevelInfo:
		prefix = color.HiBlackString(msgFmt, level.String())
	case LevelChat:
		prefix = color.CyanString(msgFmt, level.String())
	case LevelCombat:
		prefix = color.RedString(msgFmt, level.String())
	case LevelError:
		prefix = color.WhiteString(msgFmt, level.String())
	default:
		prefix = fmt.Sprintf(msgFmt, "LOG")
	}
	return prefix
}

func (l *Logger) Log(level Level, msg string) {
	prefix := selectPrefix(level)
	fmt.Fprintln(l.Writer, prefix+msg)
}

func (l *Logger) Attack(attacker string, target string, attack Attack) error {
	prefix := selectPrefix(LevelCombat)
	tpl, err := template.New("log").Parse(prefix + templateAttack + "\n")
	if err != nil {
		return err
	}

	return tpl.Execute(l.Writer, map[string]interface{}{
		"Attacker": attacker,
		"Target":   target,
		"Attack":   attack.HasName(),
		"Damage":   attack.HasDamage(),
	})
}

func (l *Logger) Status(character string, health float64, armor float64) error {
	prefix := selectPrefix(LevelInfo)
	tpl, err := template.New("log").Parse(prefix + templateStatus + "\n")
	if err != nil {
		return err
	}

	return tpl.Execute(l.Writer, map[string]interface{}{
		"Character": character,
		"Health":    health,
		"Armor":     armor,
	})
}
