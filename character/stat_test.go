package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStat(t *testing.T) {
	t.Run("new stat", func(t *testing.T) {
		arg := 100.00
		want := &Stat{value: 100.00}

		assert.Equal(t, want, NewStat(arg))
	})

	t.Run("to string", func(t *testing.T) {
		field := &Stat{value: 100.00}
                expect := "100"

		assert.Equal(t, expect, field.String())
	})

	t.Run("up", func(t *testing.T) {
		field := &Stat{value: 80.00}
		arg := 10.00
		expect := 90.00

		field.Up(arg)
		assert.Equal(t, expect, field.value)
	})

	t.Run("down", func(t *testing.T) {
		field := &Stat{value: 80.00}
		arg := 10.00
		expect := 70.00

		field.Down(arg)
		assert.Equal(t, expect, field.value)
	})

	t.Run("to value", func(t *testing.T) {
		field := &Stat{value: 80.00}
		arg := 10.00
		expect := 10.00

		field.To(arg)
		assert.Equal(t, expect, field.value)
	})
}

