package logger

import (
	"io"
	"text/template"
)

type Logger struct {
	Writer io.Writer
}

func NewLogger(writer io.Writer) *Logger {
	return &Logger{Writer: writer}
}

func (l *Logger) Attack(attacker string, target string, attack string, damage float64) error {
	tpl, err := template.New("log").Parse(`
{{.Attacker}} attacked {{.Target}} with {{.Attack}} for {{.Damage}} pts of damage
`)
	if err != nil {
		return err
	}

	return tpl.Execute(l.Writer, map[string]interface{}{
		"Attacker": attacker,
		"Target":   target,
		"Attack":   attack,
		"Damage":   damage,
	})
}

func (l *Logger) Status(character string, health string, armor string) error {
	tpl, err := template.New("log").Parse(`
Name: {{.Character}}
Health: {{.Health}}
Armor: {{.Armor}}
`)
	if err != nil {
		return err
	}

	return tpl.Execute(l.Writer, map[string]interface{}{
		"Character": character,
		"Health":    health,
		"Armor":     armor,
	})
}
