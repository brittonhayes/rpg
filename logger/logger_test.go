package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/brittonhayes/rpg/character"
	"github.com/stretchr/testify/assert"
)

func prefix(level string) string {
	return fmt.Sprintf("%-15s", strings.ToUpper(level))
}

func TestLogger(t *testing.T) {
	t.Run("new logger", func(t *testing.T) {
		arg := &bytes.Buffer{}
		expect := &Logger{Writer: &bytes.Buffer{}}

		assert.Equal(t, expect, NewLogger(arg))
	})

	t.Run("logger log info", func(t *testing.T) {
		arg := "Hello"
		field := &bytes.Buffer{}
		expect := fmt.Sprintln(prefix("INFO"), arg)

		l := NewLogger(field)
		l.Log(LevelInfo, arg)
		assert.Equal(t, expect, field.String())
	})

	t.Run("logger attack", func(t *testing.T) {
		argAttacker := "Mario"
		argTarget := "Bowser"
		argAttack := &character.Attack{
			Name:   "Stomp",
			Damage: 10.00,
		}

		field := &bytes.Buffer{}

		l := &Logger{Writer: field}
		assert.NoError(t, l.Attack(argAttacker, argTarget, argAttack))
	})

	t.Run("logger attack", func(t *testing.T) {
		argCharacter := "Mario"
		argHealth := 100.00
		argArmor := 100.00

		field := &bytes.Buffer{}

		l := &Logger{Writer: field}
		assert.NoError(t, l.Status(argCharacter, argHealth, argArmor))
	})
}
