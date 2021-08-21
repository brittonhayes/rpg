package logger

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name       string
		want       *Logger
		wantWriter string
	}{
		{name: "testing logger", want: &Logger{Writer: &bytes.Buffer{}}, wantWriter: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			assert.Equal(t, tt.want, NewLogger(writer))
			assert.Equal(t, tt.wantWriter, writer.String())
		})
	}
}

func TestLogger_Attack(t *testing.T) {
	type fields struct {
		Writer io.Writer
	}
	type args struct {
		attacker string
		target   string
		attack   string
		damage   float64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "Test attack log",
			fields: fields{
				Writer: &bytes.Buffer{},
			},
			args: args{
				attacker: "Mario",
				target:   "Bowser",
				attack:   "Stomp",
				damage:   float64(10.00),
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Writer: tt.fields.Writer,
			}
			tt.assertion(t, l.Attack(tt.args.attacker, tt.args.target, tt.args.attack, tt.args.damage))
		})
	}
}

func TestLogger_Status(t *testing.T) {
	type fields struct {
		Writer io.Writer
	}
	type args struct {
		character string
		health    string
		armor     string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "",
			fields: fields{
				Writer: &bytes.Buffer{},
			},
			args: args{
				character: "Mario",
				health:    "100",
				armor:     "100",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				Writer: tt.fields.Writer,
			}
			tt.assertion(t, l.Status(tt.args.character, tt.args.health, tt.args.armor))
		})
	}
}
