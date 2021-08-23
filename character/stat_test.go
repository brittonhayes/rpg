package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStat(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want *Stat
	}{
		{name: "Test new stat", args: args{value: 100.00}, want: &Stat{value: 100.00}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewStat(tt.args.value))
		})
	}
}

func TestStat_String(t *testing.T) {
	type fields struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Test stat to string", fields: fields{value: 100.00}, want: "100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stat{
				value: tt.fields.value,
			}
			assert.Equal(t, tt.want, s.String())
		})
	}
}

func TestStat_Up(t *testing.T) {
	type fields struct {
		value float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stat{
				value: tt.fields.value,
			}
			s.Up(tt.args.value)
		})
	}
}

func TestStat_To(t *testing.T) {
	type fields struct {
		value float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{name: "Test stat to", fields: fields{value: 30.00}, args: args{value: 50.00}, want: 50.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stat{
				value: tt.fields.value,
			}
			s.To(tt.args.value)
			assert.Equal(t, tt.want, s.value)
		})
	}
}

func TestStat_Down(t *testing.T) {
	type fields struct {
		value float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{name: "Test stat down", fields: fields{value: 20.00}, args: args{value: 10.00}, want: 10.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stat{
				value: tt.fields.value,
			}
			s.Down(tt.args.value)
			assert.Equal(t, tt.want, s.value)
		})
	}
}
